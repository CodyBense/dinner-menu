package tui

import (
	"github.com/CodyBense/dinner-menu/tui/consts"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	table table.Model
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, consts.Keymap.UpdateView):
			cmds = append(cmds, updateViewCmd())
		case key.Matches(msg, consts.Keymap.RecipeView):
			cmds = append(cmds, recipeViewCmd())
		case key.Matches(msg, consts.Keymap.Down):
			m.table.MoveDown(1)
		case key.Matches(msg, consts.Keymap.Up):
			m.table.MoveUp(1)
		case key.Matches(msg, consts.Keymap.CopyLink):
			clipboard.WriteAll(m.table.SelectedRow()[7])
		}
		switch msg.String() {
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MenuModel) View() string {
	return consts.BaseStyle.Render(m.table.View() + "\n")
}

func NewMenuModel() MenuModel {
	columns := []table.Column{
		{Title: "ID", Width: 25},
		{Title: "Name", Width: 25},
		{Title: "Cuisine", Width: 15},
		{Title: "Flavor", Width: 10},
		{Title: "Difficulty", Width: 10},
		{Title: "Time", Width: 5},
		{Title: "Liked", Width: 10},
		{Title: "Link", Width: 20},
		{Title: "Made", Width: 20},
	}

	// var rows []table.Row
	rows := []table.Row{
		{"1", "Chicken", "Italian", "Savory", "Easy", "25", "True", "https://www.recipes.com/chicken", "False"},
		{"2", "Pho", "Asian", "Savory", "Medium", "30", "True", "https://www.recipes.com/pho", "False"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := consts.SetTableStyle()
	t.SetStyles(s)

	return MenuModel{table: t}
}
