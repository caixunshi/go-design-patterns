package builder

type ColdDrink struct {
}

func (*ColdDrink) packing() Packing {
	return &Bottle{}
}
func (*ColdDrink) name() string {
	return ""
}
func (*ColdDrink) price() float64 {
	return 0
}
