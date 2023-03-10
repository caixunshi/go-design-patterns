# 单例模式
建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。

# 示例
## 场景描述
* Windows 是多进程多线程的，在操作一个文件的时候，就不可避免地出现多个进程或线程同时操作一个文件的现象，所以所有文件的处理必须通过唯一的实例来进行。
* 一些设备管理器常常设计为单例模式，比如一个电脑有两台打印机，在输出的时候就要处理不能两台打印机打印同一个文件。
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