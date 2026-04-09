package tui

import (
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
	return Model{
		step:   stepDistro,
		Distro: newSelect("distro", []string{"Debian/Ubuntu", "Rhel/Fedora", "Arch/Manjaro"}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
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
			default:
			}
		}

	}
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() tea.View {
	switch m.step {
	case stepDistro:
		return m.Distro.View()
	case stepImage:
		return m.Image.View()
	default:
		return tea.NewView(m.Distro.Value() + m.Image.Value())
	}
}

func (m Model) Value() string {
	return m.Distro.Value()
}
