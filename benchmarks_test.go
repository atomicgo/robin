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
