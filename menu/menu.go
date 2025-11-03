package menu

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	RecipeID   uint
	Name       string
	Cuisine    string
	Flavor     string
	Difficulty string
	Time       int
	Liked      bool
	Link       string
	Made       bool
}

func NewMenu(
	recipeID uint,
	name string,
	cuisine string,
	flavor string,
	difficutly string,
	time int,
	liked bool,
	link string,
	made bool,
) Menu {
	return Menu{
		RecipeID:   recipeID,
		Name:       name,
		Cuisine:    cuisine,
		Flavor:     flavor,
		Difficulty: difficutly,
		Time:       time,
		Liked:      liked,
		Link:       link,
		Made:       made,
	}
}

type GormRepository struct {
	DB *gorm.DB
}
