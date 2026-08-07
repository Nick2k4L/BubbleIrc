package models

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Nick2k4L/BubbleIrc/components"
	"github.com/Nick2k4L/BubbleIrc/theme"
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

// this eventually will inherit directly from `theme.go`

type AppModel struct {
	currentView   sessionView                 // get the currentview of the application
	channelModel  ChannelModel                // displays all of our channels within a server
	chatModel     ChatModel                   // just displays chat
	createModel   *CreateModel                // creation of our IrcClient
	inputModel    InputModel                  // input for chatting
	nameModel     NameListModel               // list of all names for a given channel
	serverModel   *ServerModel                // list of all servers
	settingsModel *SettingsModel              // change language / theme
	IrcServers    map[string]client.IRCClient // all of our IRC servers will be passed down?
	width         int
	height        int
	theme         theme.Theme
	// could potentially only live witihn
}

func InitiateApp() AppModel {
	theme := theme.BaseTheme()
	return AppModel{
		settingsModel: InitSettingsModel(),
		createModel:   InitCreateModel(theme),
		serverModel:   InitServerModel(),
		theme:         theme,
	}
}

func (m AppModel) Init() tea.Cmd {
	// always -- load theme on startUp? <-- this will become useful when start dynamically rendering
	// and possible database in the future :)
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	// extract the window width & size constantly
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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
			// TODO: this does nothing atm
		case "left":
			if m.currentView != serverView {
				m.currentView = serverView
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

	// return all batch commands, useful for initializing each screen with proper Inits()
	return m, tea.Batch(cmds...)
}

// NEED TO IMPROVE THIS AND EXTRACT OUT SOME METHODS
func (m AppModel) View() tea.View {
	var combined string
	combined = m.mainAppStyle().Render() // render our main style, rule for the terminal screen essentially

	sidebarWidth := m.width / 4 // want it to take up a quarter of the screen for the moment
	baseBoxStyle := components.BoxOptions{
		Width:                m.width - sidebarWidth,
		Height:               m.height,
		IsSubBox:             false,
		IsCenteredHorizontal: true,
	}
	sideBarStyle := components.BoxOptions{
		Width:                sidebarWidth,
		Height:               m.height,
		IsCenteredHorizontal: true,
		IsSubBox:             false,
	}

	var mainContent string
	switch m.currentView {
	case settingsView:
		mainContent = m.settingsModel.View()
	default:
		mainContent = "App model view"
	}

	// these three lines create our known screen
	mainView := components.CreateNewRoundedBoxStyle(mainContent, baseBoxStyle, m.theme)
	sideBarView := components.CreateNewRoundedBoxStyle(m.serverModel.View(), sideBarStyle, m.theme)
	combined = lipgloss.JoinHorizontal(lipgloss.Top, sideBarView, mainView) // we know we will always combine x screen with sidebar

	// we are going to return a popup box with the creation screen
	// turn into a switch for a pop up with settings screen as well!
	if m.currentView == createView {
		// Need to add a struct with properites, to help us define what the dimensions of said popup box should look like
		// Still need to do this....
		return m.returnView(components.CreatePopUpBox(combined, m.createModel.View(), m.theme))
	}

	return m.returnView(combined)
}

// just keep this true all the time
func (m AppModel) returnView(content string) tea.View {
	view := tea.NewView(content)

	view.AltScreen = true

	return view
}

// return our mainAppStyle, everything else will be under this ofc
func (m AppModel) mainAppStyle() lipgloss.Style {
	return lipgloss.NewStyle().Width(m.width).
		Height(m.height).
		Background(lipgloss.Color(m.theme.BackgroundColor)).
		Foreground(lipgloss.Color(m.theme.ForeGroundColor))
}
