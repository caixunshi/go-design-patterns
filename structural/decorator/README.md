# 装饰器模式
装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。

# 示例
## 场景描述
我们将把一个形状装饰上不同的颜色，同时又不改变形状类。
## 代码展示

* 定义画图接口
``` go
package decorator

type Shape interface {
	draw()
}
```
* 画圆形的实现
``` go
package decorator

import "fmt"

type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("这是一个圆形")
}

```
* 画三角形的实现
``` go
package decorator

import "fmt"

type Rectangle struct {
}

func (c *Rectangle) draw() {
	fmt.Println("这是一个长方形")
}

```
* 装饰器基类
``` go
package decorator

// ShapeDecorator 装饰器基类
type ShapeDecorator struct {
	Shape
}
```

* 红色装饰器
``` go
package decorator

import "fmt"

// ShapeDecorator 装饰器基类
type RedShapeDecorator struct {
    ShapeDecorator
}

func (i *RedShapeDecorator) draw() {
    i.Shape.draw()
    fmt.Println("涂上红色")
}

```

* 绿色装饰器
``` go
package decorator

import "fmt"

// GreenShapeDecorator 装饰器基类
type GreenShapeDecorator struct {
	ShapeDecorator
}

func (i *GreenShapeDecorator) draw() {
	i.Shape.draw()
	fmt.Println("涂上绿色")
}

```
* 测试类
``` go
package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	circle := &Circle{}
	rectangle := &Rectangle{}
	redDecorator1 := RedShapeDecorator{ShapeDecorator{circle}}
	redDecorator1.draw()
	fmt.Println("========")
	redDecorator2 := GreenShapeDecorator{ShapeDecorator{rectangle}}
	redDecorator2.draw()
}

```
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestDecorator
这是一个圆形
涂上红色
========
这是一个长方形
涂上绿色
--- PASS: TestDecorator (0.00s)
PASS
ok      go-design-patterns/structural/decorator 1.734s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/decorator.jpg)

## 优点
* 装饰类和被装饰类可以独立发展，不会相互耦合；
* 装饰模式是继承的一个替代模式，装饰模式可以动态扩展一个实现类的功能。 

## 缺点
* 多层装饰比较复杂。

## 使用场景
* 扩展一个类的功能。 
* 动态增加功能，动态撤销。

## 注意事项
装饰模式可以代替继承，不像继承那样耦合。