package prototype

type Rectangle struct {
	resource []interface{}
}

func (*Rectangle) getType() string {
	return "rectangle"
}
func (c *Rectangle) clone() Shape {
	return &Rectangle{
		resource: c.resource,
	}
}
