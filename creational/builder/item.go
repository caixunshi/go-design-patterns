package builder

type Item interface {
	name() string
	packing() Packing
	price() float64
}
