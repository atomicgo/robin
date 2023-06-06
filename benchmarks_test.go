package robin

import "testing"

func BenchmarkLoadbalancer_Next(b *testing.B) {
	set := []int{1, 2, 3}
	lb := NewLoadbalancer(set)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.Next()
	}
}

func BenchmarkLoadbalancer_Next_ThreadSafe(b *testing.B) {
	set := []int{1, 2, 3}
	lb := NewThreadSafeLoadbalancer(set)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.Next()
	}
}
