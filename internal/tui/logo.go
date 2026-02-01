package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const logo = `
 ▄▄▄▄▄                                           ▄▄ 
██▀▀▀▀█▄                            █▄            ██
▀██▄  ▄▀                   ▄        ██ ▄          ██
  ▀██▄▄  ▄███▀ ▄███▄ ██ ██ ████▄ ▄████ ████▄▄█▀█▄ ██
▄   ▀██▄ ██    ██ ██ ██ ██ ██ ██ ██ ██ ██   ██▄█▀ ██
▀██████▀▄▀███▄▄▀███▀▄▀██▀█▄██ ▀█▄█▀███▄█▀  ▄▀█▄▄▄▄██
`

var logoColors = [6]uint{53, 90, 127, 164, 201, 207}

var logoStyle = lipgloss.NewStyle().Margin(0, 1)

func renderLogo() string {
	lines := strings.Split(logo, "\n")
	colored_lines := make([]string, 6)
	for i, line := range lines[1 : len(lines)-1] {
		colored_lines[i] = logoStyle.Foreground(lipgloss.ANSIColor(logoColors[i])).Render(line)
	}
	return lipgloss.JoinVertical(lipgloss.Left, colored_lines...)
}
