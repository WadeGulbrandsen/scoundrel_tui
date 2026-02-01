package main

import (
	"fmt"
	"strings"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/scoundrel"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
		key.WithHelp("⏎/␣", "activate"),
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
}

func initialModel() model {
	return model{
		game:       scoundrel.New(),
		keys:       keys,
		help:       help.New(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Scoundrel")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	actions := m.game.GetActions()
	m.cursor = min(m.cursor, len(actions)-1)
	m.cursor = max(m.cursor, 0)
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

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
	actions := m.game.GetActions()

	s := "SCOUNDREL\n\n"

	s += fmt.Sprintf("Health: %d HP\n", m.game.Health)

	if m.game.Weapon == nil {
		s += "Weapon: None\n\n"
	} else {
		s += fmt.Sprintf("Weapon: %s\n", m.game.Weapon.Weapon)
		if card, err := m.game.Weapon.Killed.Peek(); err == nil {
			s += fmt.Sprintf("Last Killed: %s\n", card)
		} else {
			s += "Last Killed: None\n"
		}
	}
	s += fmt.Sprintf("Room: %v\n", m.game.Room)
	if m.game.GameOver {
		if m.game.Health > 0 {
			s += "You won! "
		} else {
			s += "You died! "
		}
		s += fmt.Sprintf("Score: %d\n\n", m.game.Score())
	} else {
		s += "Room Options:\n\n"
	}
	for i, action := range actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, action.Description)
	}
	if m.err != nil {
		s += fmt.Sprintf("Error: %v", m.err)
	}
	helpView := m.help.View(m.keys)
	height := 18 - strings.Count(s, "\n") - strings.Count(helpView, "\n")
	return "\n" + s + strings.Repeat("\n", height) + helpView
}
