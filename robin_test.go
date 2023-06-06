package robin

import (
	"fmt"
	"sync"
	"testing"
)

func TestLoadbalancer_Next(t *testing.T) {
	set := []int{1, 2, 3}

	lb := NewLoadbalancer(set)

	for i := 0; i < 10; i++ {
		if lb.Next() != set[i%len(set)] {
			t.Errorf("expected %d, got %d", set[i%len(set)], lb.Next())
		}
	}
}

func TestLoadbalancer_Next_ThreadSafe(t *testing.T) {
	var set []int

	for i := 0; i < 2000; i++ {
		set = append(set, i)
	}

	lb := NewThreadSafeLoadbalancer(set)

	var wg sync.WaitGroup

	for i := 0; i < 1337; i++ {
		wg.Add(1)
		go func() {
			lb.Next()
			wg.Done()
		}()
	}

	wg.Wait()

	if lb.Next() != 1337 {
		t.Errorf("expected %d, got %d", 1337, lb.Next())
	}
}

func TestLoadbalancer_AddItems(t *testing.T) {
	set := []int{1, 2, 3}

	lb := NewLoadbalancer(set)

	lb.AddItems(4, 5, 6)

	if lb.Items[5] != 6 {
		t.Errorf("expected %d, got %d", 6, lb.Items[5])
	}
}

func TestLoadbalancer_Reset(t *testing.T) {
	set := []int{1, 2, 3}

	lb := NewLoadbalancer(set)

	for i := 0; i < 10; i++ {
		if lb.Next() != set[i%len(set)] {
			t.Errorf("expected %d, got %d", set[i%len(set)], lb.Next())
		}
	}

	lb.Reset()

	if lb.CurrentIndex != 0 {
		t.Errorf("expected %d, got %d", 0, lb.CurrentIndex)
	}
}

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

func ExampleLoadbalancer_Next() {
	set := []int{1, 2, 3}
	lb := NewLoadbalancer(set)

	for i := 0; i < 10; i++ {
		fmt.Println(lb.Next())
	}
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
}
