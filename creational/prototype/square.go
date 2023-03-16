package prototype

type Square struct {
	resource []interface{}
}

func (*Square) getType() string {
	return "square"
}
func (s *Square) clone() Shape {
	return &Square{
		resource: s.resource,
	}
}
