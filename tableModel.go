package main

import (
	"database/sql"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codybense/dinner-menu/sqlite"
	"log"
	"strconv"
)

type TableModel struct {
	table table.Model
}

func (m TableModel) Init() tea.Cmd { return nil }

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter", "space":
			clipboard.WriteAll(m.table.SelectedRow()[6])
		case "l":
			db, err := sql.Open("sqlite", "./sqlite/recipes.db")
			if err != nil {
				log.Fatalf("Could not connect to SQLite database: %s\n", err)
			}

			defer db.Close()

			recipeName := m.table.SelectedRow()[0]
			recipeID := sqlite.GetID(db, recipeName)
			sqlite.SetLiked(db, recipeID)
			m.table.SetRows(SetRowsData())
			return m.Update(nil)
		case "u":
			input := initalModel(&m)
			return input.Update(nil)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func SetRowsData() []table.Row {
	db, err := sql.Open("sqlite", "./sqlite/recipes.db")
	if err != nil {
		log.Fatalf("Could not connect to SQLite database: %s\n", err)
	}

	defer db.Close()

	recipes, err := sqlite.FindAll(db)

	rows := []table.Row{}

	for _, recipe := range recipes {
		newRow := []table.Row{
			{
				recipe.Name,
				recipe.Cuisine_Type,
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

	return rows
}

func NewTable() TableModel {
	columns := []table.Column{
		{Title: "Name", Width: 25},
		{Title: "Cuisine Type", Width: 15},
		{Title: "Flavor", Width: 15},
		{Title: "Difficulty", Width: 10},
		{Title: "Time", Width: 10},
		{Title: "Liked", Width: 10},
		{Title: "Link", Width: 20},
		{Title: "Last Used", Width: 20},
	}

	rows := SetRowsData()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
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

	m := TableModel{t}

	return m
}
