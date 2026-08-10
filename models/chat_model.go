package models

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/Nick2k4L/BubbleIrc/theme"
)

// Every channel will have its own ChatModel
type ChatModel struct {
	// making it a string channel for now
	Inbound     chan string
	Theme       theme.Theme
	ServerModel *ServerModel

	textInput textinput.Model
}

func InitChatModel(theme theme.Theme) *ChatModel {
	t := textinput.New()
	s := t.Styles()
	t, s = theme.CreateInputBoxModelTheme(t, s)
	t.SetStyles(s)
	t.SetWidth(50)
	t.Focus()
	t.Placeholder = "Type message or /command"

	return &ChatModel{
		Inbound:   nil,
		Theme:     theme,
		textInput: t,
	}
}

func (c *ChatModel) Init() tea.Cmd {
	return textinput.Blink
}

func (c *ChatModel) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			// send message somewhere -- Eventually be the channel we are hydrating with all messages!
			// will help with bucketing and what not
			c.ServerModel.test = append(c.ServerModel.test, c.textInput.Value())
			c.textInput.Reset()
		}
	}
	c.textInput, cmd = c.textInput.Update(msg)
	return cmd
	// return nil
}

func (c *ChatModel) View() string {
	return fmt.Sprintf("This is our InputBox for Chatting: \n %s", c.textInput.View())
}
