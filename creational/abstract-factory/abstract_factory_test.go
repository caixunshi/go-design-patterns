package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
    factoryProducer := &FactoryProducer{}
    // 获取颜色工厂
    colorFactory := factoryProducer.getFactory("Color")
    redColor := colorFactory.getColor("Red")
    redColor.fill()
    greenColor := colorFactory.getColor("Green")
    greenColor.fill()
    blueColor := colorFactory.getColor("Blue")
    blueColor.fill()

    // 获取图形工厂
    shapeFactory := factoryProducer.getFactory("Shape")
    circleShape := shapeFactory.getShape("Circle")
    circleShape.draw()
    rectangleShape := shapeFactory.getShape("Rectangle")
    rectangleShape.draw()
    squareShape := shapeFactory.getShape("Square")
    squareShape.draw()
}
