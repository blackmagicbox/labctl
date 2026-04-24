package tui

import (
	"fmt"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/blackmagicbox/labctl/internal/vm"
)

type step int

const (
	stepPathToImg step = iota
	stepExistentImg
	stepDistro
	stepImage
	stepVMName
	stepHostname
	stepUsername
	stepDisk
	stepMemory
	stepCPU
	stepConfirm
	stepFinish
)

type Model struct {
	step         step
	Distro       selectModel
	Image        selectModel
	VMName       textinput.Model
	Hostname     textinput.Model
	Username     textinput.Model
	Disk         textinput.Model
	Memory       textinput.Model
	CPU          textinput.Model
	Confirmation selectModel
}

var Images = map[string][]string{
	"Debian/Ubuntu": {"ubuntu-noble24", "ubuntu-1604", "ubuntu-1804"},
	"Rhel/Fedora":   {"centos-stream-10", "fedora", "rocky-10"},
	"Arch/Manjaro":  {"steamOS", "arch-linux-basic"},
}

func New() Model {
	// VM name
	n := textinput.New()
	n.CharLimit = 156
	n.SetWidth(156)
	// Hostname
	h := textinput.New()
	h.Placeholder = "linux-lab"
	h.CharLimit = 156
	h.SetWidth(156)
	// User name
	u := textinput.New()
	u.Placeholder = "admin"
	u.CharLimit = 156
	u.SetWidth(156)
	// Disk Size
	d := textinput.New()
	d.Placeholder = "20G"
	d.CharLimit = 7
	d.SetWidth(20)
	// Memory
	m := textinput.New()
	m.Placeholder = "2048MB"
	m.CharLimit = 7
	m.SetWidth(20)
	// CPUs
	c := textinput.New()
	c.Placeholder = "2"
	c.CharLimit = 3
	c.SetWidth(20)

	return Model{
		step:         stepDistro,
		Distro:       newSelect("distro", []string{"Debian/Ubuntu", "Rhel/Fedora", "Arch/Manjaro"}),
		VMName:       n,
		Hostname:     h,
		Username:     u,
		Disk:         d,
		Memory:       m,
		CPU:          c,
		Confirmation: newSelect("Start the new VM after the wizard", []string{"Yes", "No"}),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			k := msg.String()
			if k == `q` && (m.step == stepDistro || m.step == stepImage) {
				return m, tea.Quit
			}
			switch m.step {
			case stepDistro:
				m.Distro, _ = m.Distro.Update(msg)
				if m.Distro.Chosen() {
					m.Image = newSelect("image", Images[m.Distro.Value()])
					m.step = stepImage
				}
			case stepImage:
				m.Image, _ = m.Image.Update(msg)
				if m.Image.Chosen() {
					m.VMName.Placeholder = fmt.Sprintf("%s-%s", m.Image.Value(), time.Now().Format("20060102"))
					m.VMName.Focus()
					m.step = stepVMName
				}
			case stepVMName:
				m.VMName, _ = m.VMName.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.VMName.Blur()
					if m.VMName.Value() == "" {
						m.VMName.SetValue(m.VMName.Placeholder)
					}
					m.Hostname.Focus()
					m.step = stepHostname
				}
			case stepHostname:
				m.Hostname, _ = m.Hostname.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.Hostname.Blur()
					if m.Hostname.Value() == "" {
						m.Hostname.SetValue(m.Hostname.Placeholder)
					}
					m.Username.Focus()
					m.step = stepUsername
				}
			case stepUsername:
				m.Username, _ = m.Username.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.Username.Blur()
					if m.Username.Value() == "" {
						m.Username.SetValue(m.Username.Placeholder)
					}
					m.Disk.Focus()
					m.step = stepDisk
				}
			case stepDisk:
				m.Disk, _ = m.Disk.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.Disk.Blur()
					if m.Disk.Value() == "" {
						m.Disk.SetValue(m.Disk.Placeholder)
					}
					m.Memory.Focus()
					m.step = stepMemory
				}
			case stepMemory:
				m.Memory, _ = m.Memory.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.Memory.Blur()
					if m.Memory.Value() == "" {
						m.Memory.SetValue(m.Memory.Placeholder)
					}
					m.CPU.Focus()
					m.step = stepCPU
				}
			case stepCPU:
				m.CPU, _ = m.CPU.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.CPU.Blur()
					if m.CPU.Value() == "" {
						m.CPU.SetValue(m.CPU.Placeholder)
					}
					m.step = stepConfirm
				}
			case stepConfirm:
				m.Confirmation, _ = m.Confirmation.Update(msg)
				if m.Confirmation.Chosen() {
					if m.Confirmation.Value() == "Yes" {
						// Todo: Start the new VM
					}
					m.step = stepFinish
				}
			case stepFinish:
				k := msg.String()
				if m.step > stepConfirm && k == "enter" {
					return m, tea.Quit
				}
			default:
				panic("unhandled model update")
			}

		}
	}
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() tea.View {
	out := ""
	if m.Distro.Chosen() {
		out = fmt.Sprintf("%s✔ distro: %s\n", out, m.Distro.Value())
	}
	if m.Image.Chosen() {
		out = fmt.Sprintf("%s✔ image: %s\n", out, m.Image.Value())
	}
	if m.VMName.Value() != "" && !m.VMName.Focused() {
		out = fmt.Sprintf("%s✔ VM name: %s\n", out, m.VMName.View())
	}
	if m.Hostname.Value() != "" && !m.Hostname.Focused() {
		out = fmt.Sprintf("%s✔ hostname: %s\n", out, m.Hostname.View())
	}
	if m.Username.Value() != "" && !m.Username.Focused() {
		out = fmt.Sprintf("%s✔ username: %s\n", out, m.Username.View())
	}
	if m.Disk.Value() != "" && !m.Disk.Focused() {
		out = fmt.Sprintf("%s✔ disk: %s\n", out, m.Disk.View())
	}
	if m.Memory.Value() != "" && !m.Memory.Focused() {
		out = fmt.Sprintf("%s✔ memory: %s\n", out, m.Memory.View())
	}
	if m.CPU.Value() != "" && !m.CPU.Focused() {
		out = fmt.Sprintf("%s✔ CPU: %s\n", out, m.CPU.View())
	}
	if m.Confirmation.Chosen() {
		out = fmt.Sprintf("%s✔ Start the vm: (%s)\n", out, m.Confirmation.Value())
	}

	switch m.step {
	case stepDistro:
		return tea.NewView(fmt.Sprintf("%s%v", out, m.Distro.View().Content))
	case stepImage:
		return tea.NewView(fmt.Sprintf("%s%v", out, m.Image.View().Content))
	case stepVMName:
		return tea.NewView(fmt.Sprintf("%s? VM Name: %s", out, m.VMName.View()))
	case stepHostname:
		return tea.NewView(fmt.Sprintf("%s? hostname: %s", out, m.Hostname.View()))
	case stepUsername:
		return tea.NewView(fmt.Sprintf("%s? username: %s", out, m.Username.View()))
	case stepDisk:
		return tea.NewView(fmt.Sprintf("%s? disk: %s", out, m.Disk.View()))
	case stepMemory:
		return tea.NewView(fmt.Sprintf("%s? memory: %s", out, m.Memory.View()))
	case stepCPU:
		return tea.NewView(fmt.Sprintf("%s? CPU: %s", out, m.CPU.View()))
	case stepConfirm:
		return tea.NewView(fmt.Sprintf("%s%s", out, m.Confirmation.View().Content))
	default:
		return tea.NewView(out)
	}
}

func (m Model) Value() *vm.Config {
	vmConfig := vm.Config{
		Distro:   m.VMName.Value(),
		Image:    m.Hostname.Value(),
		VMName:   m.Username.Value(),
		Username: m.Disk.Value(),
		Hostname: m.Memory.Value(),
		Disk:     m.CPU.Value(),
		Memory:   m.Distro.Value(),
		CPU:      m.Image.Value(),
	}
	return &vmConfig
}
