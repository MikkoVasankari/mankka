package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	customHelpMenu "mankka/customHelpMenu"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Creating color theme for application
var (
	White    = lipgloss.Color("#cad3f5")
	Green    = lipgloss.Color("#a6da95")
	Mauve    = lipgloss.Color("#c6a0f6")
	Lavender = lipgloss.Color("#b8c0e0")
	Blue     = lipgloss.Color("#8aadf4")
	Sapphire = lipgloss.Color("#7dc4e4")
)

type Item struct {
	title, desc string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type Model struct {
	list    list.Model
	ti      textinput.Model
	focused int
	songDir []string
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Mankka")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if m.focused == 0 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "enter", "tab":
				m.focused = 1

				cmds = m.createListOfFiles()

				return m, tea.Batch(cmds...)
			}
		}
		m.ti, cmd = m.ti.Update(msg)

	} else if m.focused == 1 {

		switch msg := msg.(type) {
		case tea.KeyMsg:

			if m.list.FilterState() == list.Filtering {
				break
			}

			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "tab":
				m.focused = 0
				m.resetLists()

			case "enter":
				var songId string

				songId = parseSongId(m.list.SelectedItem().FilterValue())

				args := []string{"--volume=50", "--geometry=800x600", "--playlist-start=" + songId, "--loop-playlist=inf"}
				args = append(args, m.songDir...)
				return m, tea.Batch(
					//tea.Printf("Selected song: %s", m.list.Index()),
					tea.ExecProcess(exec.Command("mpv", args...), nil),
				)
			}
		}
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

// Initialize styles for section
var (
	pathSelectStyle = lipgloss.NewStyle().
			Padding(0, 3).
			Foreground(Green)

	pathselectQuestionStyle = lipgloss.NewStyle().
				Padding(0, 3).
				Foreground(Blue).
				Bold(true)

	listStyle = lipgloss.NewStyle().
			Padding(0, 1).
			BorderForeground(Mauve)
)

func (m Model) View() string {

	if m.focused == 1 {
		pathSelectStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(White)
		m.ti.TextStyle = pathSelectStyle

	} else {
		pathSelectStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(Green)
	}

	return ("\n" +

		pathselectQuestionStyle.
			Render("Give a path to your directory to play files from: ") +
		"\n\n" +
		pathSelectStyle.
			Render(m.ti.View()) +
		"\n\n" +
		listStyle.
			Render(m.list.View()) +
		"\n\n" +
		fmt.Sprintln(customHelpMenu.CustomHelpView()) +
		"\n")
}

func (m *Model) resetLists() {

	if len(m.list.Items()) > 0 {
		for i := len(m.list.Items()) - 1; i >= 0; i-- {
			m.list.RemoveItem(i)
		}
	}

	m.songDir = []string{}
}

func (m *Model) createListOfFiles() []tea.Cmd {
	var (
		cmds   []tea.Cmd
		insCmd tea.Cmd
	)

	dir, _ := os.ReadDir(m.ti.Value())
	fileID := 0

	for _, value := range dir {
		if !value.IsDir() {
			m.songDir = append(m.songDir, m.ti.Value()+"/"+value.Name())
			newEntry := Item{title: strconv.Itoa(fileID) + ". " + value.Name(), desc: ""}
			insCmd = m.list.InsertItem(fileID, newEntry)
			fileID++
		}

		cmds = append(cmds, insCmd)
	}
	return cmds
}

func parseSongId(inputSong string) string {
	slicedcInputSong := strings.Split(inputSong, ".")

	return slicedcInputSong[0]
}

func initialModel() Model {

	ti := textinput.New()
	ti.Placeholder = "/Your/path/here"
	ti.Focus()
	ti.TextStyle = pathSelectStyle

	items := []list.Item{}

	listStyling := list.NewDefaultDelegate()
	listStyling.Styles.SelectedTitle = lipgloss.NewStyle().
		Foreground(Green).
		Padding(0, 2).
		Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(Green)
	listStyling.Styles.SelectedDesc = lipgloss.NewStyle().
		Foreground(Green).
		Padding(0, 2)
	listStyling.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(Lavender).
		Padding(0, 2)
	listStyling.Styles.NormalDesc = lipgloss.NewStyle().
		Foreground(Lavender).
		Padding(0, 2)

	list := list.New(items, listStyling, 100, 20)
	list.Title = "List of playable files:"
	list.Styles.Title = lipgloss.NewStyle().
		Foreground(Blue).
		Align(lipgloss.Left).
		Bold(true)
	list.Styles.NoItems = lipgloss.NewStyle().
		Padding(0, 2)
	list.SetShowStatusBar(false)
	list.SetShowHelp(false)
	//list.SetFilteringEnabled(false)

	initSongDir := []string{}

	return Model{
		list:    list,
		ti:      ti,
		focused: 0,
		songDir: initSongDir,
	}
}

func main() {

	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
