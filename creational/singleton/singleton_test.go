package singleton

import "testing"

func BenchmarkLazySingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLazySingletonInstance()
	}
}
