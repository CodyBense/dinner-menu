package tui

import (
	"github.com/CodyBense/dinner-menu/tui/consts"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type RecipeModel struct {
	table table.Model
	focus bool
}

func (m RecipeModel) Init() tea.Cmd {
	return nil
}

func (m RecipeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, consts.Keymap.UpdateView):
			cmds = append(cmds, updateViewCmd())
		case key.Matches(msg, consts.Keymap.MenuView):
			cmds = append(cmds, menuViewCmd())
		case key.Matches(msg, consts.Keymap.Down):
			m.table.MoveDown(1)
		case key.Matches(msg, consts.Keymap.Up):
			m.table.MoveUp(1)
		case key.Matches(msg, consts.Keymap.CopyLink):
			clipboard.WriteAll(m.table.SelectedRow()[7])
		}
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m RecipeModel) View() string {
	tea.Println("recipe model view refresh")
	return consts.BaseStyle.Render(m.table.View() + "\n")
}

func NewRecipeModel() RecipeModel {
	columns := []table.Column{
		{Title: "ID", Width: 25},
		{Title: "Name", Width: 25},
		{Title: "Cuisine", Width: 15},
		{Title: "Flavor", Width: 10},
		{Title: "Difficulty", Width: 10},
		{Title: "Time", Width: 5},
		{Title: "Liked", Width: 10},
		{Title: "Link", Width: 20},
		{Title: "Last Used", Width: 20},
	}

	// var rows []table.Row
	rows := []table.Row{
		{"1", "Pasta", "Italian", "Savory", "Easy", "20", "True", "https://www.recipes.com/pasta", "2025-10-31"},
		{"2", "Ramen", "Asian", "Savory", "Easy", "20", "True", "https://www.recipes.com/ramen", "2025-10-30"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := consts.SetTableStyle()
	t.SetStyles(s)

	return RecipeModel{table: t}
}
