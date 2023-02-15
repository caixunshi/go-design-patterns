package factory

import "fmt"

type Rectangle struct {
}

func (r *Rectangle) draw() {
	fmt.Println("This is a rectangle")
}
