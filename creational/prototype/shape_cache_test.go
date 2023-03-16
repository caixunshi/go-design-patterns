package prototype

import (
	"fmt"
	"testing"
)

func TestGetShape(t *testing.T) {
	shapeCache := GetShapeCache()
	rectangle := shapeCache.getShape(rectangle)
	fmt.Println(fmt.Sprintf("this is %s", rectangle.getType()))
	circle := shapeCache.getShape(circle)
	fmt.Println(fmt.Sprintf("this is %s", circle.getType()))
	square := shapeCache.getShape(square)
	fmt.Println(fmt.Sprintf("this is %s", square.getType()))
}
