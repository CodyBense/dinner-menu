package recipe

import (
	"fmt"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name       string
	Cuisine    string
	Flavor     string
	Difficulty string
	Time       int
	Liked      bool
	Link       string
	Last_Used  string
}

func NewRecipe(
	name string,
	cuisine string,
	flavor string,
	difficulty string,
	time int,
	liked bool,
	link string,
	last_used string,
) Recipe {
	return Recipe{
		Name:       name,
		Cuisine:    cuisine,
		Flavor:     flavor,
		Difficulty: difficulty,
		Time:       time,
		Liked:      liked,
		Link:       link,
		Last_Used:  last_used,
	}
}

type Repository interface {
	GetRecipeByID(recipeID uint) (Recipe, error)
	GetAllRecipes() ([]Recipe, error)
	UpdateRecipe(
		id uint,
		name string,
		cuisine string,
		flavor string,
		difficulty string,
		time int,
		liked bool,
		link string,
		last_used string,
	) error
	UpdateLiked(recipeID uint) error
}

type GormRepository struct {
	DB *gorm.DB
}

func (g *GormRepository) GetRecipeByID(recipeID uint) (Recipe, error) {
	var recipe Recipe
	if err := g.DB.Where("id = ?", recipeID).First(&recipe).Error; err != nil {
		return recipe, fmt.Errorf("Cannot find recipe: %v", err)
	}
	return recipe, nil
}

func (g *GormRepository) GetAllRecipes() ([]Recipe, error) {
	var recipes []Recipe
	if err := g.DB.Find(&recipes).Error; err != nil {
		return recipes, fmt.Errorf("Table is empty: %v", err)
	}
	return recipes, nil
}

func (g *GormRepository) UpdateRecipe(
	id uint,
	name string,
	cuisine string,
	flavor string,
	difficulty string,
	time int,
	liked bool,
	link string,
	last_used string,
) error {
	var newRecipe Recipe
	if err := g.DB.Where("id = ?", id).First(&newRecipe).Error; err != nil {
		return fmt.Errorf("Unable to update recipe: %w", err)
	}
	newRecipe.Name = name
	newRecipe.Cuisine = cuisine
	newRecipe.Flavor = flavor
	newRecipe.Difficulty = difficulty
	newRecipe.Time = time
	newRecipe.Liked = liked
	newRecipe.Link = link
	newRecipe.Last_Used = last_used
	if err := g.DB.Save(&newRecipe).Error; err != nil {
		return fmt.Errorf("Unable to save recipe: %w", err)
	}
	return nil
}

func (g *GormRepository) CreateRecipe(
	name, cuisine, flavor, difficulty string,
	time int,
	liked bool,
	link, last_used string,
) (Recipe, error) {
	recipe := Recipe{
		Name:       name,
		Cuisine:    cuisine,
		Flavor:     flavor,
		Difficulty: difficulty,
		Time:       time,
		Liked:      liked,
		Link:       link,
		Last_Used:  last_used,
	}
	if err := g.DB.Create(&recipe).Error; err != nil {
		return recipe, fmt.Errorf("Cannot create recipe: %v", err)
	}
	return recipe, nil
}

func (g *GormRepository) UpdateLiked(recipeID uint) error {
	recipe, err := g.GetRecipeByID(recipeID)
	if err != nil {
		return err
	}

	if recipe.Liked == true {
		recipe.Liked = false
	} else {
		recipe.Liked = true
	}

	if err := g.DB.Save(&recipe).Error; err != nil {
		return fmt.Errorf("Unable to save recipe: %w", err)
	}
	return nil
}
