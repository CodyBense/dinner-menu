package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codybense/dinner-menu/sqlite"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type InputModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func initalModel(tm *TableModel) InputModel {
	im := InputModel{
		inputs: make([]textinput.Model, 8),
	}

	var t textinput.Model

	for i := range im.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 50
		t.Width = 50

		switch i {
		case 0:
			t.SetValue(tm.table.SelectedRow()[0])
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.SetValue(tm.table.SelectedRow()[1])
		case 2:
			t.SetValue(tm.table.SelectedRow()[2])
		case 3:
			t.SetValue(tm.table.SelectedRow()[3])
		case 4:
			t.SetValue(tm.table.SelectedRow()[4])
		case 5:
			t.SetValue(tm.table.SelectedRow()[5])
		case 6:
			t.SetValue(tm.table.SelectedRow()[6])
		case 7:
			t.SetValue(tm.table.SelectedRow()[7])
		}

		im.inputs[i] = t
	}

	return im
}

func (im InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (im InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			tm := NewRecipeTable()
			return tm.Update(nil)
			// return im, tea.Quit

		case "ctrl+r":
			im.cursorMode++
			if im.cursorMode > cursor.CursorHide {
				im.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(im.inputs))
			for i := range im.inputs {
				cmds[i] = im.inputs[i].Cursor.SetMode(im.cursorMode)
			}

			return im, tea.Batch(cmds...)

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && im.focusIndex == len(im.inputs) {
				db, err := sql.Open("sqlite", "./sqlite/recipes.db")
				if err != nil {
					log.Fatalf("Could not connect to SQLite database: %s\n", err)
				}

				defer db.Close()
				sqlite.UpdateRecipe(db, im.inputs[0].Value(), im.inputs[1].Value(), im.inputs[2].Value(), im.inputs[3].Value(), im.inputs[4].Value(), im.inputs[5].Value(), im.inputs[6].Value(), im.inputs[7].Value())
				tm := NewRecipeTable()
				return tm.Update(nil)
				// return im, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				im.focusIndex--
			} else {
				im.focusIndex++
			}

			if im.focusIndex > len(im.inputs) {
				im.focusIndex = 0
			} else if im.focusIndex < 0 {
				im.focusIndex = len(im.inputs)
			}

			cmds := make([]tea.Cmd, len(im.inputs))
			for i := 0; i <= len(im.inputs)-1; i++ {
				if i == im.focusIndex {
					cmds[i] = im.inputs[i].Focus()
					im.inputs[i].PromptStyle = focusedStyle
					im.inputs[i].TextStyle = focusedStyle
					continue
				}

				im.inputs[i].Blur()
				im.inputs[i].PromptStyle = noStyle
				im.inputs[i].TextStyle = noStyle
			}

			return im, tea.Batch(cmds...)
		}
	}

	cmd := im.updateInputs(msg)

	return im, cmd
}

func (im *InputModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(im.inputs))

	for i := range im.inputs {
		im.inputs[i], cmds[i] = im.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (im InputModel) View() string {
	var b strings.Builder

	for i := range im.inputs {
		b.WriteString(im.inputs[i].View())
		if i < len(im.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if im.focusIndex == len(im.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(im.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

