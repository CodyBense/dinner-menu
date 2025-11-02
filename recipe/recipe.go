package recipe

type Recipe struct {
	ID         uint
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
	id uint,
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
		ID:         id,
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
