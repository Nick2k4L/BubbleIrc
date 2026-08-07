package models

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Nick2k4L/BubbleIrc/components"
	"github.com/Nick2k4L/IRC-Client/client"
)

type (
	sessionState int
	sessionView  int
)

const (
	appmodelView sessionView = iota
	channelView
	chatView
	createView
	inputView
	nameListView
	serverView
	settingsView
)

const (
	channelState sessionState = iota
	chatState
	createState
	inputState
	nameListState
	serverState
	settingsState
)

type AppModel struct {
	currentView   sessionView                 // get the currentview of the application
	channelModel  ChannelModel                // displays all of our channels within a server
	chatModel     ChatModel                   // just displays chat
	createModel   *CreateModel                // creation of our IrcClient
	inputModel    InputModel                  // input for chatting
	nameModel     NameListModel               // list of all names for a given channel
	serverModel   ServerModel                 // list of all servers
	settingsModel *SettingsModel              // change language / theme
	IrcServers    map[string]client.IRCClient // all of our IRC servers will be passed down?
	// could potentially only live witihn
}

func InitiateApp() AppModel {
	return AppModel{
		settingsModel: InitSettingsModel(),
		createModel:   InitCreateModel(),
	}
}

func (m AppModel) Init() tea.Cmd {
	// always
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	// is this a key press?
	// these are global key presses, not matter what screen we are on, we can access these with ease.
	// each screen then in the tree will have its own commands if need be
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+s":
			if m.currentView != settingsView {
				m.currentView = settingsView
				// append the init command for the spinner
				cmds = append(cmds, m.settingsModel.Init())
			} else {
				m.currentView = appmodelView
			}
		case "ctrl+a":
			if m.currentView != createView {
				m.currentView = createView
				cmds = append(cmds, m.createModel.Init())
			} else {
				m.currentView = appmodelView
			}

		}
	}

	// allows us to have our own update commands within each model.
	switch m.currentView {
	case settingsView:
		// append the update command
		cmds = append(cmds, m.settingsModel.Update(msg))
	case createView:
		cmds = append(cmds, m.createModel.Update(msg))
	case appmodelView:
	}

	// return all batch commands, useful for initializing each screen with proper
	// Inits()
	return m, tea.Batch(cmds...)
}

func (m AppModel) View() tea.View {
	// allows us to have custom views per screen
	baseBoxStyle := components.BoxOptions{
		Color:   components.Teal,
		Height:  50,
		Width:   50,
		Padding: []int{1, 1},
	}
	switch m.currentView {
	case settingsView:
		return tea.NewView(components.CreateNewRoundedBoxStyle(baseBoxStyle).Render(m.settingsModel.View()))
	case createView:
		return tea.NewView(components.CreateNewRoundedBoxStyle(baseBoxStyle).Render(m.createModel.View()))
	default:
		return tea.NewView(components.CreateNewRoundedBoxStyle(baseBoxStyle).Render("App model view"))
	}
}
