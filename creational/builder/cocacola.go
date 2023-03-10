package builder

type CocaCola struct {
	ColdDrink
}

func (*CocaCola) name() string {
	return "CocaCola"
}
func (*CocaCola) price() float64 {
	return 8.99
}
