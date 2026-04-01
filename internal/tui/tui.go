package tui

import (
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
	return textinput.Blink()
}
