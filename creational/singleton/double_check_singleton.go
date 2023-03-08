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
