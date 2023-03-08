# 单例模式
单例模式（Singleton Pattern）是 Java 中最简单的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

这种模式涉及到一个单一的类，该类负责创建自己的对象，同时确保只有单个对象被创建。这个类提供了一种访问其唯一的对象的方式，可以直接访问，不需要实例化该类的对象

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
* 在内存里只有一个实例，减少了内存的开销，尤其是频繁的创建和销毁实例；
* 避免对资源的多重占用（比如写文件操作）；
## 缺点
* 没有接口，不能继承，与单一职责原则冲突，一个类应该只关心内部逻辑，而不关心外面怎么样来实例化。

## 注意事项
getInstance() 方法中需要使用同步锁，防止多线程同时进入造成 instance 被多次实例化。