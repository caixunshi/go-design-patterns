package abstract_factory

type FactoryProducer struct {
}

func (*FactoryProducer) getFactory(choice string) AbstractFactory {
    if choice == "Color" {
        return &ColorFactory{}
    } else if choice == "Shape" {
        return &ShapeFactory{}
    } else {
        return nil
    }
}
