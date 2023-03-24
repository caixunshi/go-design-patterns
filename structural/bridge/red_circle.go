package bridge

import "fmt"

type RedCircle struct {
}

func (receiver *RedCircle) drawCircle(radius, x, y int32) {
	fmt.Println(fmt.Sprintf("draw a red circle, radius=%d, x=%d, y=%d", radius, x, y))
}
