package abstract_factory

type ShapeFactory struct {
}

func (*ShapeFactory) getShape(shape string) Shape {
    if shape == "Circle" {
        return &Circle{}
    } else if shape == "Square" {
        return &Square{}
    } else if shape == "Rectangle" {
        return &Rectangle{}
    } else {
        return nil
    }
}
func (*ShapeFactory) getColor(color string) Color {
    return nil
}
