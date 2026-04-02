package tui

import "github.com/charmbracelet/lipgloss"

var (
	labelStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	subtleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("24"))
)
