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
