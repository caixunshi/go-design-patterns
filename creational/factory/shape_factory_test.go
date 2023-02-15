package factory

import (
	"testing"
)

func TestShapeFactory(t *testing.T) {
	var factory = &ShapeFactory{}
	circle, err := factory.getShape(CircleShape)
	if err != nil {
		t.Errorf(err.Error())
	}
	circle.draw()

	rectangle, err := factory.getShape(RectangleShape)
	if err != nil {
		t.Errorf(err.Error())
	}
	rectangle.draw()

	square, err := factory.getShape(SquareShape)
	if err != nil {
		t.Errorf(err.Error())
	}
	square.draw()
}