package builder

type PepsiCola struct {
	ColdDrink
}

func (*PepsiCola) name() string {
	return "PepsiCola"
}
func (*PepsiCola) price() float64 {
	return 7.99
}
