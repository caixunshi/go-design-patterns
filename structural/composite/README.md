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
* 直接组装好树状结构返回，高层模块调用简单。

## 缺点
* 嵌套的数据结构，组装起来性能比较差。
* 在使用组合模式时，其叶子和树枝的声明都是实现类，而不是接口，违反了依赖倒置原则。

## 使用场景
您想表示对象的部分-整体层次结构（树形结构）。

## 注意事项
定义时为具体类。