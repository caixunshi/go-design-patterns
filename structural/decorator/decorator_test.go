package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	circle := &Circle{}
	rectangle := &Rectangle{}
	redDecorator1 := RedShapeDecorator{ShapeDecorator{circle}}
	redDecorator1.draw()
	fmt.Println("========")
	redDecorator2 := GreenShapeDecorator{ShapeDecorator{rectangle}}
	redDecorator2.draw()
}
