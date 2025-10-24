package main

import (
	"database/sql"
	"fmt"
	"github.com/codybense/dinner-menu/sqlite"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"os"
)

func main() {

	log_file, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Could not open log filed; %s", err)
	}

	log.SetOutput(log_file)

	db, err := sql.Open("sqlite", "./sqlite/recipes.db")
	if err != nil {
		log.Fatalf("Could not connect to SQLite database: %s", err)
	}

	defer db.Close()

	recipes, err := sqlite.FindAll(db)

	for _, recipe := range recipes {
		fmt.Println(recipe)
	}
}
