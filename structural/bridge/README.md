# 桥接模式
桥接（Bridge）是用于把抽象化与实现化解耦，使得二者可以独立变化。这种类型的设计模式属于结构型模式，它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。

这种模式涉及到一个作为桥接的接口，使得实体类的功能独立于接口实现类。这两种类型的类可被结构化改变而互不影响。

# 示例
## 场景描述
我们有一个作为桥接实现的 DrawAPI 接口和实现了 DrawAPI 接口的实体类 RedCircle、GreenCircle。Shape 是一个抽象类，将使用 DrawAPI 的对象。BridgePatternDemo 类使用 Shape 类来画出不同颜色的圆。
## 代码展示
* 定义画图接口
``` go
package bridge

type DrawAPI interface {
	drawCircle(radius, x, y int32)
}
```

* 定义画红色的圆实现
``` go
package bridge

import "fmt"

type RedCircle struct {
}

func (receiver *RedCircle) drawCircle(radius, x, y int32) {
	fmt.Println(fmt.Sprintf("draw a red circle, radius=%d, x=%d, y=%d", radius, x, y))
}
```

* 定义画绿色的圆实现
``` go
package bridge

import "fmt"

type GreenCircle struct {
}

func (g *GreenCircle) drawCircle(radius, x, y int32) {
	fmt.Println(fmt.Sprintf("draw a green circle, radius=%d, x=%d, y=%d", radius, x, y))
}
```
* 定义图形接口
``` go
package bridge

type Shape interface {
	draw()
}
```
* 定义原形的实现
``` go
package bridge

type CircleShape struct {
	DrawAPI
	radius, x, y int32
}

func (c *CircleShape) draw() {
	c.drawCircle(c.radius, c.x, c.y)
}
```
* 测试类
``` go
package bridge

import "testing"

func TestBridge(t *testing.T) {
	// draw red circle
	redCircle := &RedCircle{}
	circleShape1 := &CircleShape{
		redCircle,
		3,
		10,
		10,
	}
	circleShape1.draw()

	// draw green circle
	greenCircle := &RedCircle{}
	circleShape2 := &CircleShape{
		greenCircle,
		3,
		10,
		10,
	}
	circleShape2.draw()
}
```
可以看到，画图DrawAPI是一个独立的演进方向，而画圆形Shape又是一个独立演进的方向。
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestBridge
draw a red circle, radius=3, x=10, y=10
draw a red circle, radius=3, x=10, y=10
--- PASS: TestBridge (0.00s)
PASS
ok      go-design-patterns/structural/bridge    0.234s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/bridge.jpg)

## 优点
* 抽象和实现的分离。 
* 优秀的扩展能力。 
* 实现细节对客户透明。

## 缺点
* 桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求开发者针对抽象进行设计与编程。

## 使用场景
* 如果一个系统需要在构件的抽象化角色和具体化角色之间增加更多的灵活性，避免在两个层次之间建立静态的继承联系，通过桥接模式可以使它们在抽象层建立一个关联关系。 
* 对于那些不希望使用继承或因为多层次继承导致系统类的个数急剧增加的系统，桥接模式尤为适用。 
* 一个类存在两个独立变化的维度，且这两个维度都需要进行扩展。

## 注意事项
对于两个独立变化的维度，使用桥接模式再适合不过了.