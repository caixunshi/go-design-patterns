package decorator

import "fmt"

// GreenShapeDecorator 装饰器基类
type GreenShapeDecorator struct {
	ShapeDecorator
}

func (i *GreenShapeDecorator) draw() {
	i.Shape.draw()
	fmt.Println("涂上绿色")
}
