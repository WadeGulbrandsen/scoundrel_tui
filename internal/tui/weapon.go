package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var weaponStyle = lipgloss.NewStyle().
	Width(9).
	Height(10).
	Margin(1).
	Border(lipgloss.Border{
		Top:         "═Weapon═",
		Bottom:      "═",
		Left:        "║",
		Right:       "║",
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
	}).
	BorderForeground(gameBoarder)

func renderWeapon(m *model) string {
	weapon := ""
	if m.game.Weapon != nil {
		weapon = renderCardFront(&m.game.Weapon.Weapon)
		if card, err := m.game.Weapon.Killed.Peek(); err == nil {
			weapon = strings.Join(strings.Split(weapon, "\n")[:3], "\n")
			weapon = lipgloss.JoinVertical(lipgloss.Center, weapon, renderCardFront(&card))
		}
	}
	return weaponStyle.Render(weapon)
}
