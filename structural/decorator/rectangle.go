package decorator

import "fmt"

type Rectangle struct {
}

func (c *Rectangle) draw() {
	fmt.Println("这是一个长方形")
}
