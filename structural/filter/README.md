# 过滤器模式
过滤器模式（Filter Pattern）或标准模式（Criteria Pattern）是一种设计模式，这种模式允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来。这种类型的设计模式属于结构型模式，它结合多个标准来获得单一标准。

# 示例
## 场景描述
我们将创建一个 Person 对象、Criteria 接口和实现了该接口的实体类，来过滤 Person 对象的列表。CriteriaPatternDemo 类使用 Criteria 对象，基于各种标准和它们的结合来过滤 Person 对象的列表。

## 代码展示
* 创建Person对象
``` go
package filter

type Person struct {
	name string
	age  uint
	sex  string
}
```
* 创建过滤标准接口
``` go
package filter

type Criteria interface {
	meetCriteria([]*Person) []*Person
}
```
* 过滤男性
``` go
package filter

type MaleCriteria struct {
}

func (c *MaleCriteria) meetCriteria(persons []*Person) []*Person {
	result := make([]*Person, 0, len(persons))
	for _, person := range persons {
		if person.name == "Male" {
			result = append(result, person)
		}
	}
	return result
}
```
* 过滤女性
``` go
package filter

type FemaleCriteria struct {
}

func (c *FemaleCriteria) meetCriteria(persons []*Person) []*Person {
	result := make([]*Person, 0, len(persons))
	for _, person := range persons {
		if person.name == "Female" {
			result = append(result, person)
		}
	}
	return result
}
```
* 过滤青少年
``` go
package filter

type YoungCriteria struct {
}

func (c *YoungCriteria) meetCriteria(persons []*Person) []*Person {
	result := make([]*Person, 0, len(persons))
	for _, person := range persons {
		if person.age <= 18 {
			result = append(result, person)
		}
	}
	return result
}
```
* 组合过滤器，表示要同时满足两个条件，取交集
``` go
package filter

// AndCriteria 表示需要同时满足两个条件
type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (c *AndCriteria) meetCriteria(persons []*Person) []*Person {
	persons = c.criteria.meetCriteria(persons)
	persons = c.otherCriteria.meetCriteria(persons)
	return persons
}
```
* 组合过滤器，表示满足任意一个条件，取并集
``` go
package filter

// OrCriteria 表示取两个标准的并集
type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (c *OrCriteria) meetCriteria(persons []*Person) []*Person {
	persons1 := c.criteria.meetCriteria(persons)
	persons2 := c.otherCriteria.meetCriteria(persons)

	resultMap := make(map[string]*Person)
	for _, person := range persons1 {
		resultMap[person.name] = person
	}
	for _, person := range persons2 {
		if resultMap[person.name] == nil {
			resultMap[person.name] = person
		}
	}
	result := make([]*Person, 0, len(resultMap))
	for _, person := range resultMap {
		result = append(result, person)
	}
	return result
}
```
* 测试类
``` go
package filter

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	persons := []*Person{
		&Person{
			name: "jack",
			sex:  "Male",
			age:  25,
		},
		&Person{
			name: "marray",
			sex:  "Female",
			age:  21,
		},
		&Person{
			name: "xiaoming",
			sex:  "Male",
			age:  15,
		},
	}
	// 筛选女性
	female := &FemaleCriteria{}
	fmt.Println("开始筛选女性")
	for _, person := range female.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选女性结束")
	// 筛选女性
	male := &MaleCriteria{}
	fmt.Println("\n开始筛选男性")
	for _, person := range male.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选男性结束")
	// 筛选青少年
	young := &YoungCriteria{}
	fmt.Println("\n开始筛选青少年")
	for _, person := range young.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选青少年结束")

	// 筛选男性青少年
	and := &AndCriteria{
		criteria:      male,
		otherCriteria: young,
	}
	fmt.Println("\n开始筛选男性青少年")
	for _, person := range and.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选男性青少年结束")

	// 筛选男性青少年
	or := OrCriteria{
		criteria:      female,
		otherCriteria: young,
	}
	fmt.Println("\n开始筛选女性或青少年")
	for _, person := range or.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选女性或青少年结束")
}

```
可以看到，过滤标准可以自由组合，有新增的过滤标准时只需要新增一个类即可，这有点像责任链模式，但区别是过滤器是为了对一组对象做筛选，过滤器模式没有明确的先后关系，而责任链是有明确先后关系的，跟流水线一样，链条的每个节点只负责其中的一道工序
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestFilter
开始筛选女性
{marray 21 Female}
筛选女性结束

开始筛选男性
{jack 25 Male}
{xiaoming 15 Male}
筛选男性结束

开始筛选青少年
{xiaoming 15 Male}
筛选青少年结束

开始筛选男性青少年
{xiaoming 15 Male}
筛选男性青少年结束

开始筛选女性或青少年
{marray 21 Female}
{xiaoming 15 Male}
筛选女性或青少年结束
--- PASS: TestFilter (0.00s)
PASS
ok      go-design-patterns/structural/filter    0.009s
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