package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	id           textinput.Model
	name         textinput.Model
	cuisine_type textinput.Model
	flavor       textinput.Model
	difficulty   textinput.Model
	time         textinput.Model
	liked        textinput.Model
	link         textinput.Model
	last_used    textinput.Model
}

func (m InputModel) Init() tea.Cmd { return nil }

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m InputModel) View() string {
	return "\n"
}

func NewInput(recipes *TableModel) InputModel {
	input := InputModel{
		id: textinput.New(),
		name: textinput.New(),
		cuisine_type: textinput.New(),
		flavor: textinput.New(),
		difficulty: textinput.New(),
		time: textinput.New(),
		liked: textinput.New(),
		link: textinput.New(),
		last_used: textinput.New(),
	}

	input.id.Placeholder = recipes.table.SelectedRow()[0]
	input.name.Placeholder = recipes.table.SelectedRow()[1]
	input.cuisine_type.Placeholder = recipes.table.SelectedRow()[2]
	input.flavor.Placeholder = recipes.table.SelectedRow()[3]
	input.difficulty.Placeholder = recipes.table.SelectedRow()[4]
	input.time.Placeholder = recipes.table.SelectedRow()[5]
	input.liked.Placeholder = recipes.table.SelectedRow()[7]
	input.link.Placeholder = recipes.table.SelectedRow()[8]
	input.last_used.Placeholder = recipes.table.SelectedRow()[9]

	return input
}
