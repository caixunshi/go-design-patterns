package factory

import "errors"

const CircleShape = "circle"
const RectangleShape = "rectangle"
const SquareShape = "square"

type ShapeFactory struct {
}

func (s *ShapeFactory) getShape(shapeType string) (Shape, error) {
	switch shapeType {
	case CircleShape:
		return &Circle{}, nil
	case RectangleShape:
		return &Rectangle{}, nil
	case SquareShape:
		return &Square{}, nil
	default:
		return nil, errors.New("unknow shapeType")
	}
}