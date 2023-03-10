# 单例模式
建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。

# 示例
## 场景描述
我们假设一个快餐店的商业案例，其中，一个典型的套餐可以是一个汉堡（Burger）和一杯冷饮（Cold drink）。汉堡（Burger）可以是素食汉堡（Veg Burger）或鸡肉汉堡（Chicken Burger），它们是包在纸盒中。冷饮（Cold drink）可以是可口可乐（coke）或百事可乐（pepsi），它们是装在瓶子中。

我们将创建一个表示食物条目（比如汉堡和冷饮）的 Item 接口和实现 Item 接口的实体类，以及一个表示食物包装的 Packing 接口和实现 Packing 接口的实体类，汉堡是包在纸盒中，冷饮是装在瓶子中。

然后我们创建一个 Meal 类，带有 Item 的 ArrayList 和一个通过结合 Item 来创建不同类型的 Meal 对象的 MealBuilder。BuilderPatternDemo 类使用 MealBuilder 来创建一个 Meal。
## 代码展示
* 懒汉式
``` go
package singleton

import (
	"fmt"
	"sync"
)

var (
	lazySingletonInstance *lazySingleton
	once                  = &sync.Once{}
)

type lazySingleton struct {
}

func GetLazySingletonInstance() *lazySingleton {
	if lazySingletonInstance == nil {
		once.Do(func() {
			fmt.Println("init lazySingleton")
			lazySingletonInstance = &lazySingleton{}
		})
	}
	return lazySingletonInstance
}
```
* 懒汉式-双重锁检查
``` go
package singleton

import "sync"

type DoubleCheckSingleton struct {
}

var (
	lock                         = &sync.Mutex{}
	doubleCheckSingletonInstance *DoubleCheckSingleton
)

func GetDoubleCheckSingletonInstance() *DoubleCheckSingleton {
	if doubleCheckSingletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if doubleCheckSingletonInstance == nil {
			doubleCheckSingletonInstance = &DoubleCheckSingleton{}
		}
	}
	return doubleCheckSingletonInstance
}
```
* 饿汉式
``` go
package singleton

var hungerSingletonInstance *HungerSingleton

// init -> main
func init() {
	hungerSingletonInstance = &HungerSingleton{}
}

type HungerSingleton struct {
}

func GetHungerSingleton() *HungerSingleton {
	return hungerSingletonInstance
}
```

* 测试类
``` go
package singleton

import "testing"

func BenchmarkLazySingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instanceA := GetLazySingletonInstance()
		instanceB := GetLazySingletonInstance()
		if instanceA != instanceB {
			b.Error("different objects")
		}
	}
}

func BenchmarkHungerSingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instanceA := GetHungerSingleton()
		instanceB := GetHungerSingleton()
		if instanceA != instanceB {
			b.Error("different objects")
		}
	}
}

func BenchmarkDoubleCheckSingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instanceA := GetDoubleCheckSingletonInstance()
		instanceB := GetDoubleCheckSingletonInstance()
		if instanceA != instanceB {
			b.Error("different objects")
		}
	}
}
```
* 执行命令
```shell
go test -bench=. -run=none
```

* 输出结果
```
BenchmarkLazySingleton-2                291124850                4.460 ns/op
BenchmarkHungerSingleton-2              1000000000               0.3125 ns/op
BenchmarkDoubleCheckSingleton-2         244314862                4.992 ns/op
PASS
ok      go-design-patterns/creational/singleton 3.819s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/singleton.jpg)

## 优点
* 建造者独立，易扩展。
* 便于控制细节风险。
## 缺点
* 产品必须有共同点，范围有限制。
* 如内部变化复杂，会有很多的建造类。

## 注意事项
与工厂模式的区别是：建造者模式更加关注与零件装配的顺序。