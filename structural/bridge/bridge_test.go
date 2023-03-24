package bridge

import "testing"

func TestBridge(t *testing.T) {
	// draw red circle
	redCircle := &RedCircle{}
	circleShape1 := &CircleShape{
		redCircle,
		3,
		10,
		10,
	}
	circleShape1.draw()

	// draw green circle
	greenCircle := &RedCircle{}
	circleShape2 := &CircleShape{
		greenCircle,
		3,
		10,
		10,
	}
	circleShape2.draw()
}
