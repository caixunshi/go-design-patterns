package decorator

import "fmt"

type ShapeDecorator struct {
	shape Shape
}

func (i *ShapeDecorator) draw() {
	i.draw()
	fmt.Println("涂上颜色")
}
