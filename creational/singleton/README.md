# 单例模式
单例模式（Singleton Pattern）是 Java 中最简单的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

这种模式涉及到一个单一的类，该类负责创建自己的对象，同时确保只有单个对象被创建。这个类提供了一种访问其唯一的对象的方式，可以直接访问，不需要实例化该类的对象

# 示例
## 场景描述
* Windows 是多进程多线程的，在操作一个文件的时候，就不可避免地出现多个进程或线程同时操作一个文件的现象，所以所有文件的处理必须通过唯一的实例来进行。
* 一些设备管理器常常设计为单例模式，比如一个电脑有两台打印机，在输出的时候就要处理不能两台打印机打印同一个文件。
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
go test ./
```

* 输出结果
```
This is a circle
This is a rectangle
This is a square
--- PASS: TestShapeFactory (0.00s)
PASS
ok      go-design-patterns/creational/factory   0.330s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/singleton.jpg)

## 优点
* 在内存里只有一个实例，减少了内存的开销，尤其是频繁的创建和销毁实例；
* 避免对资源的多重占用（比如写文件操作）；
## 缺点
* 没有接口，不能继承，与单一职责原则冲突，一个类应该只关心内部逻辑，而不关心外面怎么样来实例化。

## 注意事项
getInstance() 方法中需要使用同步锁，防止多线程同时进入造成 instance 被多次实例化。