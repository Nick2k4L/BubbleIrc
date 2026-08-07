package models

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Nick2k4L/BubbleIrc/components"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle

	focusedButton = focusedStyle.Render("[ Create Server ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Create Server"))
)

type (
	creationState int
)

const (
	hostInput creationState = iota
	portInput
	nickInput
	channelInput
	tlsInput
	// TLS can we make this a radio button?
)

// this is where we will create servers
// they will then be displayed within servermodel
// each server model should have channel children
type CreateModel struct {
	index  creationState
	inputs []textinput.Model

	// this will contain all of our IRC creation logic
}

func InitCreateModel() *CreateModel {
	m := &CreateModel{
		inputs: make([]textinput.Model, 4),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		s := t.Styles()
		t.SetWidth(20)
		t.CharLimit = 32
		s.Focused.Prompt = focusedStyle
		s.Focused.Text = focusedStyle
		s.Blurred.Prompt = blurredStyle
		s.Focused.Text = focusedStyle
		s.Focused.Prompt = focusedStyle
		t.SetStyles(s)

		switch i {
		case int(hostInput):
			t.Placeholder = "irc.example.net"
			t.Focus()
		case int(portInput):
			t.Placeholder = "6697"
		case int(nickInput):
			t.Placeholder = "bubbleChat"
		case int(channelInput):
			t.SetWidth(35)
			t.Placeholder = "#lobby, #trivia (comma separated)"
		}
		m.inputs[i] = t

	}

	return m
}

func (m *CreateModel) Init() tea.Cmd {
	return textinput.Blink
}

// could have this return an IRC server + tea.cmd. We just append the tea.cmd where we need to
// could even pass in our map of irc servers
func (m *CreateModel) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "down", "enter":
			s := msg.String()
			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && int(m.index) == len(m.inputs) {
				// go through every input and update it
				// then return us to the server screen?
				// Create IRC server and jump us to Server screen
				return tea.Quit
			}

			// Cycle indexes
			if s == "up" {
				m.index--
			} else {
				m.index++
			}

			if int(m.index) > len(m.inputs) {
				m.index = hostInput
			} else if int(m.index) < 0 {
				m.index = tlsInput
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == int(m.index) {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
			}

			return tea.Batch(cmds...)
		}
	}
	cmd := m.updateInputs(msg)

	return cmd
}

func (m *CreateModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func createHeaderText(i int) string {
	switch i {
	case 0:
		return "(host)\n"
	case 1:
		return "(port)\n"
	case 2:
		return "(Nickname)\n"
	case 3:
		return "(Prejoin channels)\n"
	}
	return ""
}

func (m *CreateModel) View() string {
	var b strings.Builder
	settingsBoxOptions := components.BoxOptions{
		Color:   components.Orange,
		Height:  10,
		Width:   10,
		Padding: []int{0, 0},
	}
	for i := range m.inputs {
		b.WriteString(createHeaderText(i))
		b.WriteString(components.CreateNewRoundedBoxStyle(settingsBoxOptions).Render(m.inputs[i].View()))
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if int(m.index) == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
