package tui

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/CodyBense/dinner-menu/tui/consts"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focuedStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focuedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpState = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	focusedButton       = focuedStyle.Render("[ Submit ] ")
	blurredButton       = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type UpdateModel struct {
	focusIndex    int
	inputs        []textinput.Model
	cursorMode    cursor.Mode
	previousState sessionState
}

func (m UpdateModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m UpdateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && m.focusIndex == len(m.inputs) {
				switch m.previousState {
				case recipeView:
					m.UpdateRecipe()
				case menuView:
					m.UpdateMenu()
				}
				cmds = append(cmds, previousViewCmd())
				return m, tea.Batch(cmds...)
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmdsUpdate := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmdsUpdate[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focuedStyle
					m.inputs[i].TextStyle = focuedStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}
			cmds = append(cmds, cmdsUpdate...)
			return m, tea.Batch(cmds...)
		}
		cmds = append(cmds, cmd)
	}

	cmd = m.updateInputs(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m UpdateModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpState.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func NewUpdateModel() UpdateModel {
	m := UpdateModel{
		inputs: make([]textinput.Model, 8),
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 50
		t.Width = 50
		t.SetValue(fmt.Sprintf("%d", i))

		m.inputs[i] = t
	}

	m.inputs[0].Focus()
	m.inputs[0].PromptStyle = focuedStyle
	m.inputs[0].TextStyle = focuedStyle

	return m
}

func (m *UpdateModel) LoadInputs(mm MainModel) {
	var table table.Model

	switch mm.previousState {
	case recipeView:
		table = mm.recipeModel.table
	case menuView:
		table = mm.menuModel.table
	}

	inputs := m.inputs
	row := table.SelectedRow()

	for i := range mm.updateModel.inputs {
		switch i {
		case 0:
			inputs[0].SetValue(row[i])
			inputs[0].Focus()
			inputs[0].PromptStyle = focuedStyle
			inputs[0].TextStyle = focuedStyle
		case 1:
			inputs[1].SetValue(row[1])
		case 2:
			inputs[2].SetValue(row[2])
		case 3:
			inputs[3].SetValue(row[3])
		case 4:
			inputs[4].SetValue(row[4])
		case 5:
			inputs[5].SetValue(row[5])
		case 6:
			inputs[6].SetValue(row[6])
		case 7:
			inputs[7].SetValue(row[7])
		}
	}
}

func (m *UpdateModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m UpdateModel) UpdateRecipe() {
	var id uint
	var name, cuisine, flavor, difficulty, link, last_used string
	var time int
	var liked bool
	for i, input := range m.inputs {
		switch i {
		case 0:
			tempID, err := strconv.ParseUint(input.Value(), 10, 32)
			if err != nil {
				log.Fatalf("Couldn't parse ID: %v", err)
			}
			id = uint(tempID)
		case 1:
			name = input.Value()
		case 2:
			cuisine = input.Value()
		case 3:
			flavor = input.Value()
		case 4:
			difficulty = input.Value()
		case 5:
			tempTime, err := strconv.Atoi(input.Value())
			if err != nil {
				log.Fatalf("Couldn't parse time: %v", err)
			}
			time = tempTime
		case 6:
			tempLiked, err := strconv.ParseBool(input.Value())
			if err != nil {
				log.Fatalf("Couldn't parse time: %v", err)
			}
			liked = tempLiked
		case 7:
			link = input.Value()
		case 8:
			last_used = input.Value()
		}
	}
	consts.Rr.UpdateRecipe(id, name, cuisine, flavor, difficulty, time, liked, link, last_used)
}

func (m UpdateModel) UpdateMenu() {
	var id uint
	var name, cuisine, flavor, difficulty, link string
	var time int
	var liked, made bool
	for i, input := range m.inputs {
		switch i {
		case 0:
			tempID, err := strconv.ParseUint(input.Value(), 10, 32)
			if err != nil {
				log.Fatalf("Couldn't parse ID: %v", err)
			}
			id = uint(tempID)
		case 1:
			name = input.Value()
		case 2:
			cuisine = input.Value()
		case 3:
			flavor = input.Value()
		case 4:
			difficulty = input.Value()
		case 5:
			tempTime, err := strconv.Atoi(input.Value())
			if err != nil {
				log.Fatalf("Couldn't parse time: %v", err)
			}
			time = tempTime
		case 6:
			tempLiked, err := strconv.ParseBool(input.Value())
			if err != nil {
				log.Fatalf("Couldn't parse time: %v", err)
			}
			liked = tempLiked
		case 7:
			link = input.Value()
		case 8:
			tempMade, err := strconv.ParseBool(input.Value())
			if err != nil {
				log.Fatalf("Couldn't parse made: %v", err)
			}
			made = tempMade
		}
	}
	consts.Mr.UpdateMenu(id, name, cuisine, flavor, difficulty, time, liked, link, made)
}
