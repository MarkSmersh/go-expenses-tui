package modules

type ScreenManager struct {
	screens []Screen
	active  int
}

func NewScreenManager(screens ...Screen) ScreenManager {
	return ScreenManager{
		screens: screens,
		active:  0,
	}
}

func (m ScreenManager) GetActiveScreen() Screen {
	return m.screens[m.active]
}

func (m ScreenManager) GetActiveScreenIndex() int {
	return m.active
}

func (m *ScreenManager) SetActiveScreen(cursor int) {
	m.GetActiveScreen().SetUnactive()

	if cursor > len(m.screens)-1 {
		m.active = 0
	} else if cursor < 0 {
		m.active = len(m.screens) - 1
	} else {
		m.active = cursor
	}

	m.GetActiveScreen().SetActive()
}

func (m *ScreenManager) NextScreen() {
	m.SetActiveScreen(m.active + 1)
}

func (m *ScreenManager) PrevScreen() {
	m.SetActiveScreen(m.active - 1)
}

func (m ScreenManager) GetScreens() []Screen {
	return m.screens
}
