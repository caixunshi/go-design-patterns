# 过滤器模式
过滤器模式（Filter Pattern）或标准模式（Criteria Pattern）是一种设计模式，这种模式允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来。这种类型的设计模式属于结构型模式，它结合多个标准来获得单一标准。

# 示例
## 场景描述
我们将创建一个 Person 对象、Criteria 接口和实现了该接口的实体类，来过滤 Person 对象的列表。CriteriaPatternDemo 类使用 Criteria 对象，基于各种标准和它们的结合来过滤 Person 对象的列表。

## 代码展示
* 定义画图接口
``` go
package bridge

type DrawAPI interface {
	drawCircle(radius, x, y int32)
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
* 过滤逻辑和业务逻辑分离

## 缺点
* 桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求开发者针对抽象进行设计与编程。

## 使用场景
需要对数据进行特殊处理，处理逻辑与业务逻辑解耦的情况。

## 注意事项
在实际开发过程中，如果遇到多个过滤标准的话，可以尝试通过使用过滤器模式来将多个不同的过滤标准组合起来，使上层调用达到统一，底层筛选被隔离。