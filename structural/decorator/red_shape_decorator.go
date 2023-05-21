package decorator

import "fmt"

// ShapeDecorator 装饰器基类
type RedShapeDecorator struct {
    ShapeDecorator
}

func (i *RedShapeDecorator) draw() {
    i.Shape.draw()
    fmt.Println("涂上红色")
}
