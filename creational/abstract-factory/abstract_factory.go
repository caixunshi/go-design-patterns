package abstract_factory

type AbstractFactory interface {
    getShape(string) Shape
    getColor(string) Color
}
