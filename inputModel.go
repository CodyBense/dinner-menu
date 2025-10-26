package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// id
// name
// cuisine_type
// flavor
// difficulty
// time
// liked
// link
// last_used
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
			t.Placeholder = tm.table.SelectedRow()[0]
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = tm.table.SelectedRow()[1] 
		case 2:
			t.Placeholder = tm.table.SelectedRow()[2] 
		case 3:
			t.Placeholder = tm.table.SelectedRow()[3] 
		case 4:
			t.Placeholder = tm.table.SelectedRow()[4] 
		case 5:
			t.Placeholder = tm.table.SelectedRow()[5] 
		case 6:
			t.Placeholder = tm.table.SelectedRow()[6] 
		case 7:
			t.Placeholder = tm.table.SelectedRow()[7] 
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
			return im, tea.Quit

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
				return im, tea.Quit
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
	// var cmd tea.Cmd
	// switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	switch msg.String() {
	// 	case "esc":
	// 	case "q", "ctrl+c":
	// 		return m, tea.Quit
	// 	}
	// }
	// return m, cmd
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
	// return "\n"
}

// func NewInput(recipes *TableModel) *InputModel {
// 	input := &InputModel{
// 		id: textinput.New(),
// 		name: textinput.New(),
// 		cuisine_type: textinput.New(),
// 		flavor: textinput.New(),
// 		difficulty: textinput.New(),
// 		time: textinput.New(),
// 		liked: textinput.New(),
// 		link: textinput.New(),
// 		last_used: textinput.New(),
// 	}
//
// 	input.name.Placeholder = recipes.table.SelectedRow()[0]
// 	input.cuisine_type.Placeholder = recipes.table.SelectedRow()[1]
// 	input.flavor.Placeholder = recipes.table.SelectedRow()[2]
// 	input.difficulty.Placeholder = recipes.table.SelectedRow()[3]
// 	input.time.Placeholder = recipes.table.SelectedRow()[4]
// 	input.liked.Placeholder = recipes.table.SelectedRow()[5]
// 	input.link.Placeholder = recipes.table.SelectedRow()[6]
// 	input.last_used.Placeholder = recipes.table.SelectedRow()[7]
//
// 	input.name.Focus()
//
// 	return input
// }
