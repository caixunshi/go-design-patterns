package builder

func VegMeal() *Meal {
	meal := &Meal{}
	meal.AddItem(&VegBurger{})
	meal.AddItem(&CocaCola{})
	return meal
}

func ChickenMeal() *Meal {
	meal := &Meal{}
	meal.AddItem(&ChickenBurger{})
	meal.AddItem(&PepsiCola{})
	return meal
}
