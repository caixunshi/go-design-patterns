package abstract_factory

import "fmt"

type Shape interface {
    draw()
}

type Circle struct {
}

func (*Circle) draw() {
    fmt.Println("This is the drawing method of circle")
}

type Square struct {
}

func (*Square) draw() {
    fmt.Println("This is the drawing method of square")
}

type Rectangle struct {
}

func (*Rectangle) draw() {
    fmt.Println("This is the drawing method of rectangle")
}
