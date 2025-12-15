package screens

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/api"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/db"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/keys"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/settings"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SettingsScreen struct {
	keys  keys.KeyMapTransactionScreen
	focus modules.FocusManager
	style lipgloss.Style

	server   modules.Input
	user     modules.Input
	password modules.Input
	save     modules.Button
	reset    modules.Button
	toReset  bool

	message string

	username string
}

func NewSettingsScreen() *SettingsScreen {
	screen := SettingsScreen{
		keys: keys.TransactionScreen,
		// ipv6 adresses are included
		server:   modules.NewInput("e.x. localhost:1488", 39),
		user:     modules.NewInput("your username", 64),
		password: modules.NewInput("your password", 64).WithSecret(),
		style:    styles.Screen,
	}

	screen.save = modules.NewButton("Save", screen.saveSettings)
	screen.reset = modules.NewButton("Reset", screen.resetSettings)

	screen.setInputsFromMemory()

	screen.focus = modules.NewFocusManager(
		&screen.server,
		&screen.user,
		&screen.password,
		&screen.save,
		&screen.reset,
	)

	return &screen
}

func (s SettingsScreen) View() string {
	h := s.style.GetHeight() - 2
	w := s.style.GetWidth() - 2

	titleStyle := styles.ScreenTitle.Width(w).Height(1)

	inputStyle := lipgloss.NewStyle().Width(50).Height(2)

	centerStyle := lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Center, lipgloss.Center)

	blockStyle := lipgloss.NewStyle().
		Width(w).
		Height(h-2).
		Border(lipgloss.HiddenBorder()).
		Align(lipgloss.Center, lipgloss.Center)

	return s.style.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		titleStyle.Render("SETTINGS SCREEN"),
		blockStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				centerStyle.Render("Current user: "+s.username),
				"\n",
				inputStyle.Render(s.server.View()),
				inputStyle.Render(s.user.View()),
				inputStyle.Render(s.password.View()),
				centerStyle.Render(
					lipgloss.JoinVertical(
						lipgloss.Top,
						s.save.View()+"\n",
						s.reset.View(),
					),
				),
			),
		),
		s.message,
	))
}

func (s *SettingsScreen) Update(msg tea.Msg) modules.Cmd {
	cmd := modules.NewCmd()
	k := s.keys

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s.message = ""

		switch {
		// move tea.Quit to the global model
		case key.Matches(msg, k.Down):
			s.focus.Next()
		case key.Matches(msg, k.Up):
			s.focus.Prev()
		case key.Matches(msg, k.Activate):
			s.focus.Focused().Click()
		}

		s.focus.BlurAll()
		s.focus.Focused().Focus()
	case tea.WindowSizeMsg:
		s.style = s.style.Height(msg.Height - 2 - 2).Width(msg.Width - 2)
	}

	cmd.AddTea(
		s.server.Update(msg),
		s.user.Update(msg),
		s.password.Update(msg),
	)

	if s.toReset {
		s.toReset = false
		return cmd.WithScreen(modules.CmdAuthScreen)
	}

	return cmd
}

func (s SettingsScreen) Keys() help.KeyMap {
	return s.keys
}

func (s *SettingsScreen) SetActive() {
	logger.Logf("setting inputs from memory")
	s.setInputsFromMemory()
	s.focus.Focused().Focus()
}

func (s SettingsScreen) SetUnactive() {
	s.focus.BlurAll()
}

func (s *SettingsScreen) saveSettings() {
	// TODO: Add some regex, then handshake test with the server
	logger := modules.Logger{File: "app.log"}

	server := s.server.TextInput().Value()
	user := s.user.TextInput().Value()
	password := s.password.TextInput().Value()

	accessToken, err := api.LogIn(server, user, password)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = err.Error()
		return
	}

	err = settings.SetAccessToken(accessToken)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = "Unable to write the server (very bad error)"
		return
	}

	err = settings.SetServer(server)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = "Unable to write the server (very bad error)"
		return
	}

	err = settings.SetUsername(user)

	s.username = user

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = "Unable to write the username (very bad error)"
		return
	}

	s.password.TextInput().Reset()

	s.message = "succesfully saved and logged in to the server"
}

func (s *SettingsScreen) resetSettings() {
	db.Reset()
	s.toReset = true
}

func (s *SettingsScreen) setInputsFromMemory() {
	serverValue, err := settings.GetServer()

	logger.Logf("Server value: %s", serverValue)

	if err != nil {
		logger.Logf("%s", err.Error())
	}

	s.server.TextInput().SetValue(serverValue)

	usernameValue, err := settings.GetUsername()

	if err != nil {
		logger.Logf("%s", err.Error())
	}

	s.user.TextInput().SetValue(usernameValue)
	s.username = usernameValue
}
