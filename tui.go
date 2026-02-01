package main

import (
	"fmt"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/scoundrel"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	game   scoundrel.Game
	cursor int
	err    error
}

func initialModel() model {
	return model{
		game: scoundrel.New(),
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
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "s", "j":
			if m.cursor < len(actions)-1 {
				m.cursor++
			}
		case "enter", " ":
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

	s += fmt.Sprintf("Dungeon: %d\n", m.game.Dungeon.Length())
	s += fmt.Sprintf("Discard: %d\n", m.game.Discard.Length())
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
			s += "You won!\n"
		} else {
			s += "You died!\n"
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
	s += "\nPress q to quit.\n"
	if m.err != nil {
		s += fmt.Sprintf("Error: %v", m.err)
	}
	s += "\n\n"
	return s
}
