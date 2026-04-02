package tui

import tea "github.com/charmbracelet/bubbletea"

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

func (m selectModel) Update(msg tea.Msg) (selectModel, error) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		default:
			// no op just to shut the linter
		}
		switch msg.String() {
		case "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		default:
			// no op just to shut the linter.
		}
	}
	return m, nil
}

func (m selectModel) View() string {
	out := labelStyle.Render("? "+m.label) + ":\n"
	for i, opt := range m.options {
		if i == m.cursor {
			out += successStyle.Render(" ->" + opt + "\n")
		} else {
			out += subtleStyle.Render("  " + opt + "\n")
		}
	}

	return out
}

func (m selectModel) Value() string {
	return m.chosen
}

func (m selectModel) Chosen() bool {
	return m.chosen != ""
}
