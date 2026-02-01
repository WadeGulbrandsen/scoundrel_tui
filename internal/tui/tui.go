package tui

import (
	"time"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/scoundrel"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Help   key.Binding
	Quit   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("⏎/␣ (Enter/Spacebar)", "activate"),
	),
	Help: key.NewBinding(
		key.WithKeys("?", "h"),
		key.WithHelp("?/h", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	game       scoundrel.Game
	cursor     int
	err        error
	keys       keyMap
	help       help.Model
	inputStyle lipgloss.Style
	healthBar  progress.Model
	height     int
	width      int
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func InitialModel() model {
	return model{
		game: scoundrel.New(),
		keys: keys,
		help: help.New(),
		healthBar: progress.New(
			progress.WithDefaultGradient(),
			progress.WithoutPercentage(),
			progress.WithWidth(20),
			progress.WithFillCharacters('♥', '♡'),
		),
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	actions := m.game.GetActions()
	m.cursor = min(m.cursor, len(actions)-1)
	m.cursor = max(m.cursor, 0)
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		cmd := m.healthBar.SetPercent(float64(m.game.Health) / 20)
		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.healthBar.Update(msg)
		m.healthBar = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keys.Down):
			if m.cursor < len(actions)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.keys.Select):
			action := actions[m.cursor]
			m.err = m.game.DoAction(action)
			m.cursor = 0
		}
	}
	return m, nil
}

func (m model) View() string {
	game := renderGame(m)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, game)
}
