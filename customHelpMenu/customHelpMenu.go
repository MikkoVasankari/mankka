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
	helpText := lipgloss.NewStyle().Padding(0, 3).Render(
		fmt.Sprintf(
			"%s %s       %s  %s        %s %s       %s  %s \n%s %s    %s %s   %s %s  %s  %s \n",

			keyStyle.Render(" q "), descStyle.Render("Quit"), keyStyle.Render("Tab"), descStyle.Render("Section"),

			keyStyle.Render("↓/j"), descStyle.Render("Down"), keyStyle.Render("↑/k"), descStyle.Render("Up"),

			keyStyle.Render(" / "), descStyle.Render("Filter"), keyStyle.Render(" esc "), descStyle.Render("Reset filter"),

			keyStyle.Render("←/h"), descStyle.Render("Prev page"), keyStyle.Render("→/l"), descStyle.Render("Next page"),
		),
	)
	return helpText
}
