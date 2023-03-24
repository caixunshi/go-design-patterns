package bridge

type CircleShape struct {
	DrawAPI
	radius, x, y int32
}

func (c *CircleShape) draw() {
	c.drawCircle(c.radius, c.x, c.y)
}
