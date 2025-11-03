package consts

import (
	"github.com/CodyBense/dinner-menu/menu"
	"github.com/CodyBense/dinner-menu/recipe"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	P *tea.Program
	Rr *recipe.GormRepository
	Mr *menu.GormRepository
)

var DocStyle = lipgloss.NewStyle().Margin(0, 2)

// var BaseStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("240")).Render

type RecipeViewMsg struct{}
type MenuViewMsg struct{}
type UpdateViewMsg struct{}
type PreviousStateMsg struct{}

type keymap struct {
	UpdateView key.Binding
	RecipeView key.Binding
	MenuView   key.Binding
	Down       key.Binding
	Up         key.Binding
	CopyLink   key.Binding
	Quit       key.Binding
}

var Keymap = keymap{
	UpdateView: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "Update"),
	),
	RecipeView: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Recipe View"),
	),
	MenuView: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "Menu View"),
	),
	Down: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("j", "Down"),
	),
	Up: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("k", "Up"),
	),
	CopyLink: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Copy Link"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "Quit"),
	),
}

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func SetTableStyle() table.Styles {
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

	return s
}
