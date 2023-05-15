package decorator

import "fmt"

type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("这是一个圆形")
}
