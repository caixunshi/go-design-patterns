package builder

type Wrapper struct {
}

func (*Wrapper) pack() string {
	return "wrapper"
}
