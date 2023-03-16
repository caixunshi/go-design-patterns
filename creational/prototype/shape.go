package prototype

type Shape interface {
	getType() string
	clone() Shape
}
