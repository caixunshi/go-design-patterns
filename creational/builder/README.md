# 单例模式
建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。

# 示例
## 场景描述
我们假设一个快餐店的商业案例，其中，一个典型的套餐可以是一个汉堡（Burger）和一杯冷饮（Cold drink）。汉堡（Burger）可以是素食汉堡（Veg Burger）或鸡肉汉堡（Chicken Burger），它们是包在纸盒中。冷饮（Cold drink）可以是可口可乐（coke）或百事可乐（pepsi），它们是装在瓶子中。

我们将创建一个表示食物条目（比如汉堡和冷饮）的 Item 接口和实现 Item 接口的实体类，以及一个表示食物包装的 Packing 接口和实现 Packing 接口的实体类，汉堡是包在纸盒中，冷饮是装在瓶子中。

然后我们创建一个 Meal 类，带有 Item 的 ArrayList 和一个通过结合 Item 来创建不同类型的 Meal 对象的 MealBuilder。BuilderPatternDemo 类使用 MealBuilder 来创建一个 Meal。
## 代码展示
* 定义一个表示食物的接口
``` go
package builder

type Item interface {
	name() string
	packing() Packing
	price() float64
}
```
* 定义一个表示食物容器的接口
``` go
package builder

type Packing interface {
	pack() string
}
```
* 定义食物容器的包装袋实现
``` go
package builder

type Wrapper struct {
}

func (*Wrapper) pack() string {
	return "wrapper"
}
```
* 定义食物容器的瓶子实现
``` go
package builder

type Bottle struct {
}

func (*Bottle) pack() string {
	return "bottle"
}
```
* 定义汉堡类的食物，汉堡需要用包装袋打包
``` go
package builder

type Burger struct {
}

func (*Burger) packing() Packing {
	return &Wrapper{}
}
func (*Burger) name() string {
	return ""
}
func (*Burger) price() float64 {
	return 0
}
```

* 定义蔬菜汉堡实现
``` go
package builder

type VegBurger struct {
	Burger
}

func (*VegBurger) name() string {
	return "veg burger"
}
func (*VegBurger) price() float64 {
	return 19.99
}
```
* 定义鸡肉汉堡实现
``` go
package builder

type ChickenBurger struct {
	Burger
}

func (*ChickenBurger) name() string {
	return "chicken burger"
}
func (*ChickenBurger) price() float64 {
	return 39.99
}
```
* 定义冷饮类的食物，冷饮需要用瓶子打包
``` go
package builder

type ColdDrink struct {
}

func (*ColdDrink) packing() Packing {
	return &Bottle{}
}
func (*ColdDrink) name() string {
	return ""
}
func (*ColdDrink) price() float64 {
	return 0
}
```
* 第一种冷饮：百事可乐
``` go
package builder

type PepsiCola struct {
	ColdDrink
}

func (*PepsiCola) name() string {
	return "PepsiCola"
}
func (*PepsiCola) price() float64 {
	return 7.99
}
```
* 第二种冷饮：可口可乐
``` go
package builder

type CocaCola struct {
	ColdDrink
}

func (*CocaCola) name() string {
	return "CocaCola"
}
func (*CocaCola) price() float64 {
	return 8.99
}
```
* 定义食物套餐，套餐由一系列食物组成
``` go
package builder

import "fmt"

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	if m.items == nil {
		m.items = []Item{}
	}
	m.items = append(m.items, item)
}

func (m *Meal) ShowItems() {
	for _, item := range m.items {
		fmt.Println(fmt.Sprintf("This is a %s packed in a %s", item.name(), item.packing().pack()))
	}
}

func (m *Meal) getCost() float64 {
	money := float64(0)
	for _, item := range m.items {
		money += item.price()
	}
	return money
}
```
* 定义蔬菜汉堡套餐建造者和鸡肉汉堡套餐构建者
``` go
package builder

func VegMeal() *Meal {
	meal := &Meal{}
	meal.AddItem(&VegBurger{})
	meal.AddItem(&CocaCola{})
	return meal
}

func ChickenMeal() *Meal {
	meal := &Meal{}
	meal.AddItem(&ChickenBurger{})
	meal.AddItem(&PepsiCola{})
	return meal
}
```
* 定义测试类
``` go
package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	vegMeal := VegMeal()
	fmt.Println("veg meal item:")
	vegMeal.ShowItems()
	fmt.Println(fmt.Sprintf("veg meal，total cost：%f", vegMeal.getCost()))

	chickenMeal := ChickenMeal()
	fmt.Println("chicken meal item:")
	chickenMeal.ShowItems()
	fmt.Println(fmt.Sprintf("chicken meal，total cost：%f", chickenMeal.getCost()))
}
```

```
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestBuilder
veg meal item:
This is a veg burger packed in a wrapper
This is a CocaCola packed in a bottle
veg meal，total cost：28.980000
chicken meal item:
This is a chicken burger packed in a wrapper
This is a PepsiCola packed in a bottle
chicken meal，total cost：47.980000
--- PASS: TestBuilder (0.00s)
PASS
ok      go-design-patterns/creational/builder   0.339s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/builder.jpg)

## 优点
* 建造者独立，易扩展。
* 便于控制细节风险。
## 缺点
* 产品必须有共同点，范围有限制。
* 如内部变化复杂，会有很多的建造类。

## 注意事项
与工厂模式的区别是：建造者模式更加关注与零件装配的顺序。