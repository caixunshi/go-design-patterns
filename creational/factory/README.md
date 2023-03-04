# 工厂模式
工厂模式（Factory Pattern）是 Java 中最常用的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。

# 示例
## 场景描述
我需要一个画图的工具，分别可以画长方形，原型，正方形，我不需要知道如何创建这个画图工具，只需要告诉图形工厂
我需要画什么图形，工厂就可以把这个画图工具返回给我。
## 代码展示
* 画图接口
``` go
type Shape interface {
	draw()
}
```
* 不同图形的实现
``` go
// 长方形
type Rectangle struct {
}
func (r *Rectangle) draw() {
	fmt.Println("This is a rectangle")
}

// 圆型
type Circle struct {
}
func (c *Circle) draw() {
	fmt.Println("This is a circle")
}

// 正方形
type Square struct {
}
func (s *Square) draw() {
	fmt.Println("This is a square")
}
```
* 图形工厂
``` go
const CircleShape = "circle"
const RectangleShape = "rectangle"
const SquareShape = "square"

type ShapeFactory struct {
}

func (s *ShapeFactory) getShape(shapeType string) (Shape, error) {
	switch shapeType {
	case CircleShape:
		return &Circle{}, nil
	case RectangleShape:
		return &Rectangle{}, nil
	case SquareShape:
		return &Square{}, nil
	default:
		return nil, errors.New("unknow shapeType")
	}
}
```
* 测试类
``` go
func TestShapeFactory(t *testing.T) {
	var factory = &ShapeFactory{}
	circle, err := factory.getShape(CircleShape)
	if err != nil {
		t.Errorf(err.Error())
	}
	circle.draw()

	rectangle, err := factory.getShape(RectangleShape)
	if err != nil {
		t.Errorf(err.Error())
	}
	rectangle.draw()

	square, err := factory.getShape(SquareShape)
	if err != nil {
		t.Errorf(err.Error())
	}
	square.draw()
}
```
* 执行命令
```shell
go test -v shape_factory_test.go shape_factory.go shape.go square.go circle.go rectangle.go
```

* 输出结果
```
=== RUN   TestShapeFactory
This is a circle
This is a rectangle
This is a square
--- PASS: TestShapeFactory (0.00s)
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/factory.jpg)

## 优点
* 一个调用者想创建一个对象，只要知道其名称就可以了；
* 扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以；
* 屏蔽产品的具体实现，调用者只关心产品的接口；
## 缺点
* 每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖，这并不是什么好事。

## 注意事项
作为一种创建类模式，在任何需要生成复杂对象的地方，都可以使用工厂方法模式。有一点需要注意的地方就是复杂对象适合使用工厂模式，而简单对象，特别是只需要通过 new 就可以完成创建的对象，
无需使用工厂模式。如果使用工厂模式，就需要引入一个工厂类，会增加系统的复杂度。