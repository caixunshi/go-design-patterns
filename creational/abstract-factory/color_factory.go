package abstract_factory

type ColorFactory struct {
}

func (*ColorFactory) getColor(color string) Color {
    if color == "Red" {
        return &Red{}
    } else if color == "Green" {
        return &Green{}
    } else if color == "Blue" {
        return &Blue{}
    } else {
        return nil
    }
}

func (*ColorFactory) getShape(shape string) Shape {
    return nil
}
