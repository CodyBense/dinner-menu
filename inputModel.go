package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputModel struct {
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

func (m inputModel) Init() tea.Cmd { return nil }

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m inputModel) View() string {
	return "\n"
}
