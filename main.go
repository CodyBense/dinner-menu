package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CodyBense/dinner-menu/tui"
	"github.com/CodyBense/dinner-menu/tui/consts"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if f, err := tea.LogToFile("logs/debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging: ", err)
		err := os.MkdirAll("./logs", 0755)
		if err != nil {
			fmt.Println("Couldn't create the logd dir: ", err)
			panic(err)
		}
		os.Create("./logs/debug.log")
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	m, _ := tui.InitModel()

	consts.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := consts.P.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
