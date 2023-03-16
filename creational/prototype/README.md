# 原型模式
原型模式（Prototype Pattern）适用于创建重复的对象，减少创建对象过程中的耗时操作，同时又能保证性能。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式之一。

# 示例
## 场景描述
我们将创建一个抽象类 Shape 和扩展了 Shape 类的实体类。下一步是定义类 ShapeCache，该类把 shape 对象存储在一个 Hashtable 中，并在请求的时候返回它们的克隆。

PrototypePatternDemo 类使用 ShapeCache 类来获取 Shape 对象。
## 代码展示
* 图形抽象
``` go
package prototype

type Shape interface {
	getType() string
	clone() Shape
}
```
* Rectangle实现
``` go
package prototype

type Rectangle struct {
	resource []interface{}
}

func (*Rectangle) getType() string {
	return "rectangle"
}
func (c *Rectangle) clone() Shape {
	return &Rectangle{
		resource: c.resource,
	}
}
```
* Circle实现
``` go
package prototype

type Circle struct {
	resource []interface{}
}

func (*Circle) getType() string {
	return "circle"
}

func (c *Circle) clone() Shape {
	return &Circle{
		resource: c.resource,
	}
}
```
* Square实现
``` go
package prototype

type Square struct {
	resource []interface{}
}

func (*Square) getType() string {
	return "square"
}
func (c *Square) clone() Shape {
	return &Square{
		resource: c.resource,
	}
}
```
* Cache实现
``` go
package prototype

import "sync"

type ShapeCache struct {
	cache map[ShapeType]Shape
}

var (
	shapeCache *ShapeCache
	once       sync.Once
)

type ShapeType string

const (
	rectangle ShapeType = "rectangle"
	circle              = "circle"
	square              = "square"
)

func GetShapeCache() *ShapeCache {
	if shapeCache == nil {
		once.Do(func() {
			cache := make(map[ShapeType]Shape)
			cache[rectangle] = &Rectangle{
				resource: []interface{}{},
			}
			cache[circle] = &Circle{
				resource: []interface{}{},
			}
			cache[square] = &Square{
				resource: []interface{}{},
			}
			shapeCache = &ShapeCache{
				cache: cache,
			}

		})
	}
	return shapeCache
}

func (s *ShapeCache) getShape(shapeType ShapeType) Shape {
	return s.cache[shapeType].clone()
}
```
* 测试类
``` go
package prototype

import (
	"fmt"
	"testing"
)

func TestGetShape(t *testing.T) {
	shapeCache := GetShapeCache()
	rectangle := shapeCache.getShape(rectangle)
	fmt.Println(fmt.Sprintf("this is %s", rectangle.getType()))
	circle := shapeCache.getShape(circle)
	fmt.Println(fmt.Sprintf("this is %s", circle.getType()))
	square := shapeCache.getShape(square)
	fmt.Println(fmt.Sprintf("this is %s", square.getType()))
}
```
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestGetShape
this is rectangle
this is circle
this is square
--- PASS: TestGetShape (0.00s)
PASS
ok      go-design-patterns/creational/prototype 0.429s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/prototype.jpg)

## 优点
* 性能提高。
* 逃避构造函数的约束。
## 缺点
* 在go中需要自己实现一个Clone方法，要注意共享变量带来的线程安全问题。

## 注意事项
这种模式在go中一般使用对象池模式，所谓对象池，就是创建一个容器，将需要的对象提前创建出来并缓存起来，需要用时，直接从对象池中获取，通过这种方式既可以提升获取对象的速度，又可以控制全局对象的数量。
