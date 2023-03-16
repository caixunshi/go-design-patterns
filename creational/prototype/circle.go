package prototype

type Circle struct {
	resource []interface{}
}

func (*Circle) getType() string {
	return "circle"
}

func (c *Circle) clone() Shape {
	return &Circle{
		resource: c.resource,
	}
}
