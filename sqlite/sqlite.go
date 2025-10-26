package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codybense/dinner-menu/recipe"
	_ "github.com/glebarez/go-sqlite"
)

type Menu struct {
	Recipe recipe.Recipe
	Made   bool
}

func FindAll(db *sql.DB) ([]recipe.Recipe, error) {
	query := "SELECT * FROM recipes"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Find all query failed: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	var recipes []recipe.Recipe

	for rows.Next() {
		r := &recipe.Recipe{}
		err := rows.Scan(&r.Id, &r.Name, &r.Cuisine_Type, &r.Flavor, &r.Difficulty, &r.Time, &r.Liked, &r.Link, &r.Last_Used)
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

	var recipes []recipe.Recipe

	for rows.Next() {
		r := &recipe.Recipe{}
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

	var recipes []recipe.Recipe

	for rows.Next() {
		r := &recipe.Recipe{}
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

func UpdateRecipe(db *sql.DB, name, cuisine_type, flavor, difficulty, time, liked, link, last_used string) {
	recipeID := GetID(db, name)

	updateQuery := fmt.Sprintf(`UPDATE recipes
	SET name = '%s', 
		cuisine_type = '%s',
		flavor = '%s',
		difficulty = '%s',
		time = '%s',
		liked = '%s',
		link = '%s',
		last_used = '%s'
	WHERE
		id = %d 
	`, name, cuisine_type, flavor, difficulty, time, liked, link, last_used, recipeID)

	_, err := db.Exec(updateQuery)
	if err != nil {
		log.Fatalf("UpdatedRecipe query failed: %s", err)
	}
}
