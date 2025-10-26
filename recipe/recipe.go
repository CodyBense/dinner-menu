package recipe

type Recipe struct {
	Id           int
	Name         string
	Cuisine_Type string
	Flavor       string
	Difficulty   string
	Time         int
	Liked        bool
	Link         string
	Last_Used    string
}

func NewRecipe( id int, name, cuisin_type, flavor, difficulty string, time int, liked bool, link string, last_used string) Recipe {
	recipe := Recipe{
		Id: id,
		Name: name,
		Cuisine_Type: cuisin_type,
		Flavor: flavor,
		Difficulty: difficulty,
		Time: time,
		Liked: liked,
		Link: link,
		Last_Used: last_used,
	}

	return recipe
}
