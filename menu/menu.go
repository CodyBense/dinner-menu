package menu

import (
	"fmt"

	"gorm.io/gorm"
)

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

type Repository interface {
	GetMenuByID(menuID uint) (Menu, error)
	GetAllMenu() ([]Menu, error)
	UpdateMenu(id, recipeID uint, name, cuisine, flavor, difficulty string, time int, liked bool, link string, made bool) error
	CreateMenu(id, recipeID uint, name, cuisine, flavor, difficulty string, time int, liked bool, link string, made bool) (Menu, error)
	UpdateMade(menuID uint) error
	RemoveMenu(menuID uint) error
}

type GormRepository struct {
	DB *gorm.DB
}

func (g *GormRepository) GetMenuByID(menuiD uint) (Menu, error) {
	var menu Menu
	if err := g.DB.Where("id = ?", menuiD).First(&menu).Error; err != nil {
		return menu, fmt.Errorf("Cannot find menu item: %v", err)
	}
	return menu, nil
}

func (g *GormRepository) GetAllMenu() ([]Menu, error) {
	var menus []Menu
	if err := g.DB.Find(&menus).Error; err != nil {
		return menus, fmt.Errorf("Table is empty: %v", err)
	}
	return menus, nil
}

func (g *GormRepository) UpdateMenu(
	id uint,
	name, cuisine, flavor, difficulty string,
	time int,
	liked bool,
	link string,
	made bool,
) error {
	var newMenu Menu
	if err := g.DB.Where("id = ?", id).First(&newMenu).Error; err != nil {
		return fmt.Errorf("Unable to update Menu: %v", err)
	}

	newMenu.Name = name
	newMenu.Cuisine = cuisine
	newMenu.Flavor = flavor
	newMenu.Difficulty = difficulty
	newMenu.Time = time
	newMenu.Liked = liked
	newMenu.Link = link
	newMenu.Made = made

	if err := g.DB.Save(&newMenu).Error; err != nil {
		return fmt.Errorf("Unable to save menu: %v", err)
	}
	return nil
}

func (g *GormRepository) CreateMenu(
	recipeID uint,
	name, cuisine, flavor, difficulty string,
	time int,
	liked bool,
	link string,
	made bool,
) (Menu, error) {
	menu := Menu{
		RecipeID:   recipeID,
		Name:       name,
		Cuisine:    cuisine,
		Flavor:     flavor,
		Difficulty: difficulty,
		Time:       time,
		Liked:      liked,
		Link:       link,
		Made:       made,
	}

	if err := g.DB.Create(&menu).Error; err != nil {
		return menu, fmt.Errorf("Couldn't create menu: %v", err)
	}
	return menu, nil
}

func (g *GormRepository) UpdateMade(menuID uint) error {
	menu, err := g.GetMenuByID(menuID)
	if err != nil {
		return err
	}

	if menu.Made == false {
		menu.Made = true
	} else {
		menu.Made = false
	}

	if err := g.DB.Save(&menu).Error; err != nil {
		return fmt.Errorf("Couldn't save menu: %v", err)
	}
	return nil
}

func (g *GormRepository) RemoveMenu(menuID uint) error {
	result := g.DB.Delete(&Menu{}, menuID)
	return result.Error
}
