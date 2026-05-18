package cloudconfig

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/blackmagicbox/labctl/internal/vm"
)

var templates embed.FS

func Generate(cfg *vm.Config) (string, error) {
	// Resolve and create the output directory
	dir, err := outputDir(cfg.VMName)
	if err != nil {
		return "", fmt.Errorf("preparing output directory: %w", err)
	}

	// Iterate over a map of ´templateName.tmpl -> outpuTname` (no extension)
	files := map[string]string{
		"user-data.tmpl": "user-data",
		"meta-data.tmpl": "meta-data",
	}
	// For each template render the file and save it to disk.
	for tmplName, outputName := range files {
		content, err := renderTemplate(tmplName, *cfg)
		if err != nil {
			return "", fmt.Errorf("rendering template %s: %w", tmplName, err)
		}

		outPath := filepath.Join(dir, outputName)

		if err := os.WriteFile(outPath, content, 0644); err != nil {
			return "", fmt.Errorf("writing file %s: %w", outPath, err)
		}
	}
	// On success return the directory path to cmd/new.go
	return dir, nil
}

func outputDir(vmName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving home directory: %w", err)
	}
	dir := filepath.Join(home, ".local", "share", "labctl", "vms", vmName)
	if err := os.Mkdir(dir, 0755); err != nil {
		return "", fmt.Errorf("creating output directory: %w", err)
	}
	return dir, nil
}

func renderTemplate(name string, cfg vm.Config) ([]byte, error) {
	content, err := templates.ReadFile(fmt.Sprintf("templates/%s", name))
	if err != nil {
		return nil, fmt.Errorf("reading template: %w", err)
	}
	tmpl, err := template.New(name).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, cfg); err != nil {
		return nil, fmt.Errorf("executing template %s: %w", name, err)
	}
	return buf.Bytes(), nil
}
