package builder

type ChickenBurger struct {
	Burger
}

func (*ChickenBurger) name() string {
	return "chicken burger"
}
func (*ChickenBurger) price() float64 {
	return 39.99
}
