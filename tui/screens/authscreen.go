package screens

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/api"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/keys"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/settings"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthScreen struct {
	keys  keys.KeyMapAuthScreen
	focus modules.FocusManager

	server   modules.Input
	user     modules.Input
	password modules.Input
	login    modules.Button
	register modules.Button

	isAuth bool

	message string
}

func NewAuthScreen() *AuthScreen {
	screen := AuthScreen{
		keys: keys.AuthScreen,
		// ipv6 adresses are included
		server:   modules.NewInput("e.x. localhost:1488", 39),
		user:     modules.NewInput("your username", 64),
		password: modules.NewInput("your password", 64).WithSecret(),
	}

	screen.login = modules.NewButton("Login", screen.logIn)
	screen.register = modules.NewButton("Register", screen.signUp)

	screen.focus = modules.CreateFocusManager(
		&screen.server,
		&screen.user,
		&screen.password,
		&screen.login,
		&screen.register,
	)

	return &screen
}

func (s AuthScreen) View() string {
	view := `                                                                                   
   mmm   mmmm         mmmmmm m    m mmmmm  mmmmmm mm   m  mmmm  mmmmmm  mmmm        
 m"   " m"  "m        #       #  #  #   "# #      #"m  # #"   " #      #"   "       
 #   mm #    #        #mmmmm   ##   #mmm#" #mmmmm # #m # "#mmm  #mmmmm "#mmm        
 #    # #    #        #       m""m  #      #      #  # #     "# #          "#       
  "mmm"  #mm#         #mmmmm m"  "m #      #mmmmm #   ## "mmm#" #mmmmm "mmm#"       
    `

	view += "\n\n" + s.server.View()

	view += "\n\n" + s.user.View()

	view += "\n\n" + s.password.View()

	view += "\n\n" + s.login.View()

	view += "\n\n" + s.register.View()

	if len(s.message) > 0 {
		view += "\n\n" + s.message
	}

	return view
}

func (s *AuthScreen) Update(msg tea.Msg) modules.Cmd {
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
		case key.Matches(msg, k.Quit):
			return cmd.WithTea(tea.Quit)
		}

		s.focus.BlurAll()
		s.focus.Focused().Focus()
	}

	if s.isAuth {
		s.isAuth = false
		return cmd.WithScreen(modules.CmdExclusiveOff)
	}

	return cmd
}

func (s AuthScreen) Keys() help.KeyMap {
	return s.keys
}

func (s *AuthScreen) GetTextInputs() []*textinput.Model {
	tis := []*textinput.Model{
		s.server.TextInput(),
		s.user.TextInput(),
		s.password.TextInput(),
	}

	return tis
}

func (s AuthScreen) SetActive() {
	s.focus.Focused().Focus()
}

func (s AuthScreen) SetUnactive() {
	s.focus.BlurAll()
}

// TODO: Add some regex, then handshake test with the server
func (s *AuthScreen) logIn() {
	server := s.server.TextInput().Value()
	user := s.user.TextInput().Value()
	password := s.password.TextInput().Value()

	accessToken, err := api.LogIn(server, user, password)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = err.Error()
		return
	}

	err = s.saveCredentials(accessToken, server, user)

	if err != nil {
		return
	}

	s.isAuth = true

	s.message = "Succesfully logged in!"
}

func (s *AuthScreen) signUp() {
	server := s.server.TextInput().Value()
	user := s.user.TextInput().Value()
	password := s.password.TextInput().Value()

	accessToken, err := api.SignUp(server, user, password)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = err.Error()
		return
	}

	err = s.saveCredentials(accessToken, server, user)

	if err != nil {
		return
	}

	s.isAuth = true

	s.message = "Succesfully registered!"
}

func (s *AuthScreen) saveCredentials(accessToken string, server string, user string) error {
	err := settings.SetAccessToken(accessToken)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = "Unable to write the server (very bad error)"
		return err
	}

	err = settings.SetServer(server)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = "Unable to write the server (very bad error)"
		return err
	}

	err = settings.SetUsername(user)

	if err != nil {
		logger.Logf("%s", err.Error())
		s.message = "Unable to write the username (very bad error)"
		return err
	}

	return nil
}
