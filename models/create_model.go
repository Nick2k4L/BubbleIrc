package models

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Nick2k4L/BubbleIrc/components"
	"github.com/Nick2k4L/BubbleIrc/theme"
)

var (
// focusedStyle = lipgloss.NewStyle().Foreground(color.Black)
// blurredStyle = lipgloss.NewStyle().Foreground(color.White)
// focusedButton = focusedStyle.Render("[ Create Server ]")
// blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Create Server"))
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
	theme  theme.Theme
	// this will contain all of our IRC creation logic --Reroute somwhere
}

// want to make this a pop up instead
// when click ctrl+a we should see a popup window
func InitCreateModel(theme theme.Theme) *CreateModel {
	m := &CreateModel{
		inputs: make([]textinput.Model, 4),
		theme:  theme,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()

		s := t.Styles()
		t.SetWidth(20)
		t.CharLimit = 32

		// Prompt color i.e background 'char' of character we end up choosing
		s.Focused.Prompt = lipgloss.NewStyle().Foreground((lipgloss.Color(theme.FocusedStyle))).Background(lipgloss.Color(theme.BackgroundColor))
		s.Blurred.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.UnFocusedStyle)).Background(lipgloss.Color(theme.BackgroundColor))

		// Text colors
		s.Focused.Text = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.FocusedStyle)).Background(lipgloss.Color(theme.BackgroundColor))
		s.Blurred.Text = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.UnFocusedStyle)).Background(lipgloss.Color(theme.BackgroundColor))

		// Color for blinking cursor
		s.Cursor.Color = lipgloss.Color(theme.UnFocusedStyle)

		// background color for placeholder
		s.Focused.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.UnFocusedStyle)).Background(lipgloss.Color(theme.BackgroundColor))
		s.Blurred.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.UnFocusedStyle)).Background(lipgloss.Color(theme.BackgroundColor))

		t.Prompt = "* " // Can use this to create custom prompt '>' indicators later...

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
			t.Placeholder = "#lobby, #trivia"
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

// View model for the create, everything
func (m *CreateModel) View() string {
	var b strings.Builder

	// Subbox we are wrapping it in
	settingsBoxOptions := components.BoxOptions{
		Width:    30,
		IsSubBox: true,
	}

	for i := range m.inputs {
		b.WriteString(createHeaderText(i))
		b.WriteString(components.CreateNewRoundedBoxStyle(m.inputs[i].View(), settingsBoxOptions, m.theme))
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := buttonCreation("Create Server", m.theme, false)
	if int(m.index) == len(m.inputs) {
		button = buttonCreation("---> Create Server <---", m.theme, true)
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", button)

	return b.String()
}

// decided to take this out into its own function from the tutorial
func buttonCreation(text string, theme theme.Theme, focused bool) string {
	if focused {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(theme.FocusedStyle)).Render(text)
	}
	x := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.UnFocusedStyle)).Render(text)
	// return fmt.Sprintf("[ %s ]", lipgloss.NewStyle().Foreground(lipgloss.Color(theme.UnFocusedStyle)).Render(text))
	return x
}
