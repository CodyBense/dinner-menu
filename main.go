package main

import (
	"fmt"
	"log"

	"github.com/CodyBense/dinner-menu/menu"
	"github.com/CodyBense/dinner-menu/recipe"
	"github.com/CodyBense/dinner-menu/tui"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./sql/recipes.db"), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("Unable to open sqlite database: %w", err)
	}

	err = db.AutoMigrate(&recipe.Recipe{}, &menu.Menu{})

	return db, nil
}

func main() {
	db, err := OpenSqlite()
	if err != nil {
		log.Fatal(err)
	}

	rr := recipe.GormRepository{DB: db}
	mr := menu.GormRepository{DB: db}

	tui.StartTea(rr, mr)
	// if f, err := tea.LogToFile("logs/debug.log", "help"); err != nil {
	// 	fmt.Println("Couldn't open a file for logging: ", err)
	// 	err := os.MkdirAll("./logs", 0755)
	// 	if err != nil {
	// 		fmt.Println("Couldn't create the log dir: ", err)
	// 		panic(err)
	// 	}
	// 	os.Create("./logs/debug.log")
	// 	os.Exit(1)
	// } else {
	// 	defer func() {
	// 		err = f.Close()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}()
	// }
	//
	// db, err := OpenSqlite()
	// if err != nil {
	// 	log.Fatalf("Couldn't open sqlite database: %w", err )
	// }
	//
	// rr := recipe.GormRepository{DB: db}
	// mr := menu.GormRepository{DB: db}
	//
	// m, _ := tui.InitModel()
	//
	// consts.P = tea.NewProgram(m, tea.WithAltScreen())
	// if _, err := consts.P.Run(); err != nil {
	// 	fmt.Println("Error running program: ", err)
	// 	os.Exit(1)
	// }
}
