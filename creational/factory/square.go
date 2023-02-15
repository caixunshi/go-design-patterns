package factory

import "fmt"

type Square struct {
}

func (s *Square) draw() {
	fmt.Println("This is a square")
}