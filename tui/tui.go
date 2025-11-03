package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/CodyBense/dinner-menu/menu"
	"github.com/CodyBense/dinner-menu/recipe"
	"github.com/CodyBense/dinner-menu/tui/consts"
	tea "github.com/charmbracelet/bubbletea"
)

func StartTea(rr recipe.GormRepository, mr menu.GormRepository) error {
	if f, err := tea.LogToFile("logs/debug.log", "help "); err != nil {
		fmt.Println("Couldn't open a file for logging: ", err)
		err := os.MkdirAll("./logs", 0755)
		if err != nil {
			fmt.Println("Couldn't create the log dir: ", err)
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

	consts.Rr = &rr
	consts.Mr = &mr

	// m, err := InitModel()
	m, _ := InitModel()
	// if err != nil {
	// 	log.Fatalf("Failed to init Model: %v", err)
	// }

	consts.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := consts.P.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}

	return nil
}
