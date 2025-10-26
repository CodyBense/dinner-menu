package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"os"
)

type status int

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var modles []tea.Model

const (
	recipes_table status = iota
	menu_table
	update_text
)

func main() {
	log_file, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Could not open log filed; %s\n", err)
	}

	log.SetOutput(log_file)

	m := InitTableModel()
	// if _, err := tea.NewProgram(m).Run(); err != nil {
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatalf("Error running program: %s\n", err)
	}
}
