package abstract_factory

import "fmt"

type Color interface {
    fill()
}

type Red struct {
}

func (*Red) fill() {
    fmt.Println("fill in red")
}

type Green struct {
}

func (*Green) fill() {
    fmt.Println("fill in green")
}

type Blue struct {
}

func (*Blue) fill() {
    fmt.Println("fill in blue")
}
