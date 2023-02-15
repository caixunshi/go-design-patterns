package factory

import "fmt"

type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("This is a circle")
}