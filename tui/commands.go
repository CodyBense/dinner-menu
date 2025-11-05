package tui

import (
	"github.com/CodyBense/dinner-menu/tui/consts"
	tea "github.com/charmbracelet/bubbletea"
)

func menuViewCmd() tea.Cmd {
	return func() tea.Msg {
		return consts.MenuViewMsg{}
	}
}

func recipeViewCmd() tea.Cmd {
	return func() tea.Msg {
		return consts.RecipeViewMsg{}
	}
}

func updateViewCmd() tea.Cmd {
	return func() tea.Msg {
		return consts.UpdateViewMsg{}
	}
}

func previousViewCmd() tea.Cmd {
	return func() tea.Msg {
		return consts.PreviousStateMsg{}
	}
}

func addToMenuCmd() tea.Cmd {
	return func() tea.Msg {
		return consts.AddToMenuMsg{}
	}
}

func removeFromMenuCMD() tea.Cmd {
	return func() tea.Msg {
		return consts.RemoveFromMenuMsg{}
	}
}
