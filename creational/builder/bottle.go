package builder

type Bottle struct {
}

func (*Bottle) pack() string {
	return "bottle"
}
