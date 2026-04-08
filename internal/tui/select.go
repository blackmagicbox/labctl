package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type selectModel struct {
	label   string
	options []string
	cursor  int
	chosen  string
}

func newSelect(label string, options []string) selectModel {
	return selectModel{
		label:   label,
		options: options,
	}
}

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) Update(msg tea.Msg) (selectModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			m.chosen = m.options[m.cursor]
		}
	}
	return m, nil
}

func (m selectModel) View() tea.View {
	out := labelStyle.Render("? "+m.label) + ":\n"
	for i, opt := range m.options {
		if i == m.cursor {
			out += successStyle.Render("-> "+opt) + "\n"
		} else {
			out += subtleStyle.Render("   "+opt) + "\n"
		}
	}
	if m.Chosen() {
		return tea.NewView(successStyle.Render(fmt.Sprintf("✓ %s: %s", m.label, m.Value())) + "\n")
	}

	return tea.NewView(out)
}

func (m selectModel) Value() string {
	return m.chosen
}

func (m selectModel) Chosen() bool {
	return m.chosen != ""
}
