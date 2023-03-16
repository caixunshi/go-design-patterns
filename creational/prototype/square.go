package prototype

type Square struct {
	resource []interface{}
}

func (*Square) getType() string {
	return "square"
}
func (c *Square) clone() Shape {
	return &Square{
		resource: c.resource,
	}
}
