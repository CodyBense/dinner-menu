package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

type Recipe struct {
	Id          int
	Name        string
	Cusine_Type string
	Flavor      string
	Difficulty  string
	Time        int
	Liked       bool
	Link        string
	Last_Used   string
}

func FindAll(db *sql.DB) ([]Recipe, error) {
	sql := "SELECT * FROM recipes"

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalf("Find all query failed: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		r := &Recipe{}
		err := rows.Scan(&r.Id, &r.Name, &r.Cusine_Type, &r.Flavor, &r.Difficulty, &r.Time, &r.Liked, &r.Link, &r.Last_Used)
		if err != nil {
			log.Fatalf("Database row scan failed: %s\n", err)
			return nil, err
		}
		recipes = append(recipes, *r)
	}

	return recipes, nil
}
