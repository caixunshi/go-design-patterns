# 抽象工厂模式
抽象工厂模式（Abstract Factory Pattern）是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。
# 示例
## 场景描述
工作了，为了参加一些聚会，肯定有两套或多套衣服吧，比如说有商务装（成套，一系列具体产品）、时尚装（成套，一系列具体产品），甚至对于一个家庭来说，可能有商务女装、商务男装、时尚女装、时尚男装，这些也都是成套的，即一系列具体产品。假设一种情况（现实中是不存在的，要不然，没法进入共产主义了，但有利于说明抽象工厂模式），在您的家中，某一个衣柜（具体工厂）只能存放某一种这样的衣服（成套，一系列具体产品），每次拿这种成套的衣服时也自然要从这个衣柜中取出了。用 OOP 的思想去理解，所有的衣柜（具体工厂）都是衣柜类的（抽象工厂）某一个，而每一件成套的衣服又包括具体的上衣（某一具体产品），裤子（某一具体产品），这些具体的上衣其实也都是上衣（抽象产品），具体的裤子也都是裤子（另一个抽象产品）。

## 代码展示
* 图形接口
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

func (*Rectangle) draw() {
    fmt.Println("This is the drawing method of rectangle")
}

// 圆型
type Circle struct {
}

func (*Circle) draw() {
    fmt.Println("This is the drawing method of circle")
}

// 正方形
type Square struct {
}

func (*Square) draw() {
    fmt.Println("This is the drawing method of square")
}
```
* 定义图形工厂
``` go
type ShapeFactory struct {
}

func (*ShapeFactory) getShape(shape string) Shape {
    if shape == "Circle" {
        return &Circle{}
    } else if shape == "Square" {
        return &Square{}
    } else if shape == "Rectangle" {
        return &Rectangle{}
    } else {
        return nil
    }
}
func (*ShapeFactory) getColor(color string) Color {
    return nil
}
```
* 定义颜色接口
``` go
type Color interface {
    fill()
}
```
* 不同颜色的实现
``` go
// 红色
type Red struct {
}

func (*Red) fill() {
    fmt.Println("fill in red")
}

// 绿色
type Green struct {
}

func (*Green) fill() {
    fmt.Println("fill in green")
}

// 蓝色
type Blue struct {
}

func (*Blue) fill() {
    fmt.Println("fill in blue")
}
```
* 定义抽象工厂
``` go
type AbstractFactory interface {
    getShape(string) Shape
    getColor(string) Color
}
```
* 定义抽象工厂创造类
``` go
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
```
* 测试类
``` go
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
```
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestAbstractFactory
fill in red
fill in green
fill in blue
This is the drawing method of circle
This is the drawing method of rectangle
This is the drawing method of square
--- PASS: TestAbstractFactory (0.00s)
PASS
ok      go-design-patterns/creational/abstract-factory  0.095s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/abstract-factory.jpg)

## 优点
当一个产品族中的多个对象被设计成一起工作时，它能保证客户端始终只使用同一个产品族中的对象。
## 缺点
产品族扩展非常困难，要增加一个系列的某一产品，既要在抽象的 Creator 里加代码，又要在具体的里面加代码。
## 注意事项
要在抽象工厂增加一类新的工厂比较麻烦，但是要更换已有的工厂比较简单。