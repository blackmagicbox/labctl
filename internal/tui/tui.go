package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var labelStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

type Model struct {
	input     textinput.Model
	submitted bool
	err       error
}

func New(defaultname string) Model {
	ti := textinput.New()
	ti.Placeholder = defaultname
	ti.Focus()
	ti.CharLimit = 64
	return Model{input: ti}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd

}

func (m Model) View() string {
	return fmt.Sprintf("%s %s\n", labelStyle.Render("? VM name:"), m.input.View())
}

func (m Model) Value() string {
	return m.input.Value()
}
