package builder

type Burger struct {
}

func (*Burger) packing() Packing {
	return &Wrapper{}
}
func (*Burger) name() string {
	return ""
}
func (*Burger) price() float64 {
	return 0
}
