package entity


type Sweet struct {
	ID		int	   `json:"id"`
	Title	string `json:"title"`
	Type	string `json:"type"`
	Kcal	int	   `json:"kcal"`
	Weight  string `json:"height"`
	Price	int    `json:"price"`
}


var Sweets = []Sweet{

	{
		ID: 1,
		Title: "Mochi",
		Type:"Japanese rice cake",
		Weight: "100 gr.",
		Kcal: 379,
		Price: 400,
	},
	{
		ID: 2,
		Title: "Baklava",
		Type:"traditional Turkish pastry",
		Weight: "100 gr.",
		Kcal: 433,
		Price: 583,
	},
	{
		ID: 3,
		Title: "Eclair",
		Type:"traditional French pastry",
		Weight: "100 gr.",
		Kcal: 414,
		Price: 325,
	},
	{
		ID: 4,
		Title: "Barmbrack",
		Type:"Sweet irish bread with raisin",
		Weight: "100 gr.",
		Kcal: 262,
		Price: 250,
	},
	{
		ID: 5,
		Title: "Cheesecake",
		Type:"international dessert, consisting of cream cheese and cookies",
		Weight: "100 gr.",
		Kcal: 276,
		Price: 320,
	},
}
