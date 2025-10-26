package main

type Recipe struct {
	Name        string
	Cusine_Type string
	Flavor      string
	Difficulty  string
	Time        int
	Liked       bool
	Link        string
	Last_Used   string
}

func NewRecipe() Recipe {
}
