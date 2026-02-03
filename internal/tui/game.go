package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	// gameBackground = lipgloss.ANSIColor(234)
	gameBoarder = lipgloss.ANSIColor(127)
)

var gameStyle = lipgloss.NewStyle().
	BorderForeground(gameBoarder)

func renderGame(m model) string {
	logo := renderLogo()
	dungeon := renderDeck(&m.game.Dungeon)
	discard := renderDeck(&m.game.Discard)
	health := fmt.Sprintf("\nHealth: %3d HP ", m.game.Health) + m.healthBar.View()
	status := "\n"
	switch {
	case m.game.GameOver && m.game.Health > 0:
		status += fmt.Sprintf("You won! Score: %d", m.game.Score())
	case m.game.GameOver:
		status += fmt.Sprintf("You lost! Score: %d", m.game.Score())
	case m.err != nil:
		status += fmt.Sprintf("Error: %v", m.err)
	}
	header := lipgloss.JoinHorizontal(lipgloss.Top,
		dungeon,
		gameStyle.Render(lipgloss.JoinVertical(lipgloss.Center,
			gameStyle.Render(logo),
			gameStyle.Render(health),
			gameStyle.Render(status),
		)),
		discard,
	)
	weapon := renderWeapon(&m)
	room := renderRoom(m)
	middle := lipgloss.JoinHorizontal(lipgloss.Top, weapon, room)
	body := lipgloss.JoinVertical(lipgloss.Center, header, middle)
	helpView := m.help.View(m.keys)
	height := 34 - strings.Count(body, "\n") - strings.Count(helpView, "\n")
	game := body + strings.Repeat("\n", height) + helpView
	return gameStyle.Render(game)
}

func renderSlot(m model, i int) string {
	content := ""
	if m.game.Room[i] != nil {
		content = renderCardFront(m.game.Room[i])
	}
	top := lipgloss.PlaceHorizontal(9, lipgloss.Center, strconv.Itoa((i + 1)), lipgloss.WithWhitespaceChars("═"))
	boarder := lipgloss.Border{
		Top:         top,
		Bottom:      "═",
		Left:        "║",
		Right:       "║",
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
	}
	return deckStyle.Border(boarder).Render(content)
}

func renderRoom(m model) string {
	slots := []string{}
	for i := range m.game.Room {
		slots = append(slots, renderSlot(m, i))
	}
	actions := renderActions(m)
	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Center, slots...),
		actions,
	)
}
