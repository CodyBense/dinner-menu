package tui

import (
	"github.com/CodyBense/dinner-menu/tui/consts"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	recipeView sessionState = iota
	menuView
	updateView
)

type MainModel struct {
	state         sessionState
	previousState sessionState
	recipeModel   RecipeModel
	menuModel     MenuModel
	updateModel   UpdateModel
	// rr          *recipe.GormRepository
	// mr          *menu.GormRepository
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case consts.RecipeViewMsg:
		m.previousState = m.state
		m.state = recipeView
	case consts.MenuViewMsg:
		m.previousState = m.state
		m.state = menuView
	case consts.UpdateViewMsg:
		m.previousState = m.state
		m.updateModel.LoadInputs(m)
		m.state = updateView
	case consts.PreviousStateMsg:
		m.state = m.previousState
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, consts.Keymap.Quit):
			return m, tea.Quit
		}
	}

	switch m.state {
	case recipeView:
		newRecipe, newCmd := m.recipeModel.Update(msg)
		recipeModel, ok := newRecipe.(RecipeModel)
		if !ok {
			panic("failed assertion onto recipe model")
		}
		m.recipeModel = recipeModel
		cmd = newCmd
	case menuView:
		newMenu, newCmd := m.menuModel.Update(msg)
		menuModel, ok := newMenu.(MenuModel)
		if !ok {
			panic("failed assertion onto menu model")
		}
		m.menuModel = menuModel
		cmd = newCmd
	case updateView:
		newUpdate, newCmd := m.updateModel.Update(msg)
		updateModel, ok := newUpdate.(UpdateModel)
		if !ok {
			panic("failed assertion onto update model")
		}
		m.updateModel = updateModel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case menuView:
		return m.menuModel.View()
	case updateView:
		return m.updateModel.View()
	default:
		m.recipeModel.table.SetRows(SetRecipeRowData())
		return m.recipeModel.View()
	}
}

func InitModel() (tea.Model, tea.Cmd) {
	m := MainModel{state: recipeView, recipeModel: NewRecipeModel(), menuModel: NewMenuModel(), updateModel: NewUpdateModel()}

	return m, func() tea.Msg { return nil }
}
