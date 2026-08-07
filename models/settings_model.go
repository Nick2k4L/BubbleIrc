package models

import (
	"fmt"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SettingsModel struct {
	text    string
	visited int
	spinner spinner.Model
}

func InitSettingsModel() *SettingsModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000"))
	return &SettingsModel{
		text:    "In settings model",
		visited: 0,
		spinner: s,
	}
}

func (m *SettingsModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *SettingsModel) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	// is this a key press?
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+d":
			return tea.Quit
		case "ctrl+s":
			m.visited++

		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return cmd
	}

	return nil
}

func (m *SettingsModel) View() string {
	return fmt.Sprintf("%s + %d\n %s", m.text, m.visited, m.spinner.View())
}
