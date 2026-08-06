package models

import tea "charm.land/bubbletea/v2"

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
	currentView   sessionView
	channelModel  ChannelModel
	chatModel     ChatModel
	createModel   CreateModel
	inputModel    InputModel
	nameModel     NameListModel
	serverModel   ServerModel
	settingsModel SettingsModel
}

func InitiateApp() AppModel {
	return AppModel{}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// is this a key press?
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+s":
			if m.currentView == appmodelView {
				m.currentView = settingsView
			} else {
				m.currentView = appmodelView
			}

		}
	}
	return m, nil
}

func (m AppModel) View() tea.View {
	switch m.currentView {
	case settingsView:
		return tea.NewView(m.settingsModel.View())
	default:
		return tea.NewView("In the App model")
	}
}
