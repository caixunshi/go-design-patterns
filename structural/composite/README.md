# 组合模式
组合模式（Composite Pattern），又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。组合模式依据树形结构来组合对象，用来表示部分以及整体层次。这种类型的设计模式属于结构型模式，它创建了对象组的树形结构。

# 示例
## 场景描述
演示了一个组织中员工的层次结构。

## 代码展示
* 创建Employee对象
``` go
package composite

type Employee struct {
	name         string
	dept         string
	salary       float32
	subordinates []Employee
}

```

* 测试类
``` go
package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	ceo := Employee{"John", "CEO", 30000, nil}

	// 总裁办公室
	headSales := Employee{"Robert", "Head Sales", 20000, nil}
	headMarketing := Employee{"Michel", "Head Marketing", 20000, nil}
	headSubordinates := []Employee{headSales, headMarketing}
	ceo.subordinates = headSubordinates

	// 文员
	clerk1 := Employee{"Laura", "Marketing", 10000, nil}
	clerk2 := Employee{"Bob", "Marketing", 10000, nil}
	marketingSubordinates := []Employee{clerk1, clerk2}
	headMarketing.subordinates = marketingSubordinates

	// 销售
	salesExecutive1 := Employee{"Richard", "Sales", 10000, nil}
	salesExecutive2 := Employee{"Rob", "Sales", 10000, nil}
	salesSubordinates := []Employee{salesExecutive1, salesExecutive2}
	headSales.subordinates = salesSubordinates

	//打印该组织的所有员工
	fmt.Println(ceo)
	fmt.Println(&ceo.subordinates[1] == &headMarketing)
	fmt.Println(headMarketing.subordinates)
}

```
可以看到组合模式创建了一个包含自己对象组的类，非常适合用来表示自包含的数据结构，如组织架构，地址信息等
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
![类图](https://caixunshi.github.io/document/go-design-patterns/cpmposite.jpg)

## 优点
* 将对象的过滤、校验逻辑抽离出来，降低系统的复杂度。
* 过滤规则可实现重复利用。

## 缺点
* 性能较低，每个过滤器对每一个元素都会进行遍历。如果有n个元素，m个过滤器，则复杂度为O(mn)

## 使用场景
需要对数据进行特殊处理，处理逻辑与业务逻辑解耦的情况。

## 注意事项
在实际开发过程中，如果遇到多个过滤标准的话，可以尝试通过使用过滤器模式来将多个不同的过滤标准组合起来，使上层调用达到统一，底层筛选被隔离。