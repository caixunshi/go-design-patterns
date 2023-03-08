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
