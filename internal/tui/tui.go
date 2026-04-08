package tui

import (
	tea "charm.land/bubbletea/v2"
)

type step int

const (
	stepDistro step = iota
)

type Model struct {
	step           step
	Distro         selectModel
	SelectedDistro string
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
			m.Distro, _ = m.Distro.Update(msg)
			m.SelectedDistro = m.Distro.Value()
		}
	}
	var cmd tea.Cmd
	return m, cmd

}

func (m Model) View() tea.View {
	return m.Distro.View()
}

func (m Model) Value() string {
	return m.SelectedDistro
}
