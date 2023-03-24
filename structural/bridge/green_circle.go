package bridge

import "fmt"

type GreenCircle struct {
}

func (g *GreenCircle) drawCircle(radius, x, y int32) {
	fmt.Println(fmt.Sprintf("draw a green circle, radius=%d, x=%d, y=%d", radius, x, y))
}
