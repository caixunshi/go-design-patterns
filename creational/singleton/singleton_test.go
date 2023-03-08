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
