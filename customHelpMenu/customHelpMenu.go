package customHelpMenu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	Blue = lipgloss.Color("#8aadf4")

	keyStyle  = lipgloss.NewStyle().Foreground(Blue).Bold(true)
	descStyle = lipgloss.NewStyle().Foreground(Blue)
)

func CustomHelpView() string {
	helpText := lipgloss.NewStyle().Width(50).Padding(0, 3).Render(
		fmt.Sprintf(
			"%s %s       %s  %s \n%s %s       %s  %s \n%s %s    %s %s \n%s %s  %s  %s \n",

			keyStyle.Render(" q "), descStyle.Render("quit"),
			keyStyle.Render("Tab"), descStyle.Render("switch section"),

			keyStyle.Render("↓/j"), descStyle.Render("down"),
			keyStyle.Render("↑/k"), descStyle.Render("up"),

			keyStyle.Render(" / "), descStyle.Render("filter"),
			keyStyle.Render(" esc "), descStyle.Render("reset filtering"),

			keyStyle.Render("←/h"), descStyle.Render("prev page"),
			keyStyle.Render("→/l"), descStyle.Render("next page"),
		),
	)
	return helpText
}
