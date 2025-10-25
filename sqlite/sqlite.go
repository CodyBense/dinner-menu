package sqlite

import (
	"database/sql"
	"fmt"
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

type Menu struct {
	Recipe Recipe
	Made   bool
}

func FindAll(db *sql.DB) ([]Recipe, error) {
	query := "SELECT * FROM recipes"

	rows, err := db.Query(query)
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
			log.Fatalf("FindAll row scan failed: %s\n", err)
			return nil, err
		}
		recipes = append(recipes, *r)
	}

	return recipes, nil
}

func GetID(db *sql.DB, recipeName string) int {
	query := fmt.Sprintf("SELECT id FROM recipes WHERE name = '%s'", recipeName)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("GetID query failed: %s\n", err)
		return 0
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		r := &Recipe{}
		err := rows.Scan(&r.Id)
		if err != nil {
			log.Fatalf("GetID row scan failed: %s\n", err)
		}
		recipes = append(recipes, *r)
	}

	return recipes[0].Id
}

func GetLiked(db *sql.DB, recipeID int) bool {
	getQuery := fmt.Sprintf("SELECT liked FROM recipes WHERE id = '%d'", recipeID)

	rows, err := db.Query(getQuery)
	if err != nil {
		log.Fatalf("Upate getQuery failed: %s\n", err)
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		r := &Recipe{}
		err := rows.Scan(&r.Liked)
		if err != nil {
			log.Fatalf("Upate getQuery row scan failed: %s\n", err)
		}

		recipes = append(recipes, *r)
	}

	return recipes[0].Liked
}

func SetLiked(db *sql.DB, recipeID int) {

	recipeLiked := GetLiked(db, recipeID)

	if recipeLiked == false {
		updateQuery := fmt.Sprintf("UPDATE recipes SET liked = 1 WHERE id = %d", recipeID)

		_, err := db.Exec(updateQuery)
		if err != nil {
			log.Fatalf("UpdateLiked query failed: %s\n", err)
		}
	} else {
		updateQuery := fmt.Sprintf("UPDATE recipes SET liked = 0 WHERE id = %d", recipeID)

		_, err := db.Exec(updateQuery)
		if err != nil {
			log.Fatalf("UpdateLiked query failed: %s\n", err)
		}
	}

}
