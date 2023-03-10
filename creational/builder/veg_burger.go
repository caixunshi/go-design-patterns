package builder

type VegBurger struct {
	Burger
}

func (*VegBurger) name() string {
	return "veg burger"
}
func (*VegBurger) price() float64 {
	return 19.99
}
