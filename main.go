package main

import (
	"database/sql"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codybense/dinner-menu/sqlite"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"os"
	"strconv"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
	log_file, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Could not open log filed; %s\n", err)
	}

	log.SetOutput(log_file)

	db, err := sql.Open("sqlite", "./sqlite/recipes.db")
	if err != nil {
		log.Fatalf("Could not connect to SQLite database: %s\n", err)
	}

	defer db.Close()

	recipes, err := sqlite.FindAll(db)

	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Cuisine Type", Width: 20},
		{Title: "Flavor", Width: 20},
		{Title: "Difficulty", Width: 20},
		{Title: "Time", Width: 20},
		{Title: "Liked", Width: 20},
		{Title: "Link", Width: 20},
		{Title: "Last Used", Width: 20},
	}

	rows := []table.Row{}

	for _, recipe := range recipes {
		newRow := []table.Row{
			{
				recipe.Name,
				recipe.Cusine_Type,
				recipe.Flavor,
				recipe.Difficulty,
				strconv.Itoa(recipe.Time),
				strconv.FormatBool(recipe.Liked),
				recipe.Link,
				recipe.Last_Used,
			},
		}

		rows = append(rows, newRow...)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)


	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatalf("Error running program: %s\n", err)
	}
}
