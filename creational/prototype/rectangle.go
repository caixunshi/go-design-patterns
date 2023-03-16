package prototype

type Rectangle struct {
	resource []interface{}
}

func (*Rectangle) getType() string {
	return "rectangle"
}
func (r *Rectangle) clone() Shape {
	return &Rectangle{
		resource: r.resource,
	}
}
