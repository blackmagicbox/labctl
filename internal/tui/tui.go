package tui

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
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
)

type Model struct {
	step     step
	Distro   selectModel
	Image    selectModel
	VMName   textinput.Model
	Hostname textinput.Model
	Username textinput.Model
}

var Images = map[string][]string{
	"Debian/Ubuntu": {"ubuntu-noble24", "ubuntu-1604", "ubuntu-1804"},
	"Rhel/Fedora":   {"centos-stream-10", "fedora", "rocky-10"},
	"Arch/Manjaro":  {"steamOS", "arch-linux-basic"},
}

func New() Model {
	n := textinput.New()
	n.Placeholder = "example: ubuntu-noble"
	n.CharLimit = 156
	n.SetWidth(20)
	h := textinput.New()
	h.Placeholder = "example: ubuntu-1604"
	h.CharLimit = 156
	h.SetWidth(20)

	return Model{
		step:     stepDistro,
		Distro:   newSelect("distro", []string{"Debian/Ubuntu", "Rhel/Fedora", "Arch/Manjaro"}),
		VMName:   n,
		Hostname: h,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			switch m.step {
			case stepDistro:
				m.Distro, _ = m.Distro.Update(msg)
				if m.Distro.Chosen() {
					m.Image = newSelect("image:", Images[m.Distro.Value()])
					m.step = stepImage
				}
			case stepImage:
				m.Image, _ = m.Image.Update(msg)
				if m.Image.Chosen() {
					m.step = stepVMName
				}
			case stepVMName:
				m.VMName.Focus()
				m.VMName, _ = m.VMName.Update(msg)
				if msg.Code == tea.KeyEnter {
					m.VMName.SetValue(m.VMName.Value() + "\n")
					m.step = stepHostname
				}
			default:
			}
		}

	}
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() tea.View {
	label := ""
	switch m.step {
	case stepDistro:
		return m.Distro.View()
	case stepImage:
		label += fmt.Sprintf("✔ distro: %s\n", m.Distro.View().Content)
		return tea.NewView(fmt.Sprintf("%s %s \n", label, m.Image.View().Content))
	case stepVMName:
		return tea.NewView(m.VMName.View())
	default:
		return tea.NewView(m.Distro.Value() + m.Image.Value())
	}
}

func (m Model) Value() string {
	return m.Distro.Value()
}
