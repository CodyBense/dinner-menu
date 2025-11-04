package tui

import (
	"log"
	"strconv"

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
		case key.Matches(msg, consts.Keymap.UpdateMade):
			tempID, err := strconv.ParseUint(m.table.SelectedRow()[0], 10, 32)
			if err != nil {
				log.Fatalf("Couldn't parse ID; %v", err)
			}
			id := uint(tempID)
			consts.Mr.UpdateMade(id)
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
		{Title: "ID", Width: 5},
		{Title: "Name", Width: 25},
		{Title: "Cuisine", Width: 15},
		{Title: "Flavor", Width: 10},
		{Title: "Difficulty", Width: 10},
		{Title: "Time", Width: 5},
		{Title: "Liked", Width: 10},
		{Title: "Link", Width: 20},
		{Title: "Made", Width: 20},
	}

	rows := SetMenuRowData()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := consts.SetTableStyle()
	t.SetStyles(s)

	return MenuModel{table: t}
}

func SetMenuRowData() []table.Row {
	var rows []table.Row
	menus, err := consts.Mr.GetAllMenu()
	if err != nil {
		log.Fatal(err)
	}

	for _, menu := range menus {
		newRow := table.Row{
			strconv.Itoa(int(menu.ID)),
			menu.Name,
			menu.Cuisine,
			menu.Flavor,
			menu.Difficulty,
			strconv.Itoa(menu.Time),
			strconv.FormatBool(menu.Liked),
			menu.Link,
			strconv.FormatBool(menu.Made),
		}
		rows = append(rows, newRow)
	}

	return rows
}
