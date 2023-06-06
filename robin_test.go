package robin

import (
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

	lbs := []Loadbalancer[int]{*NewLoadbalancer(set), *NewThreadSafeLoadbalancer(set)}

	for lbi := range lbs {
		lbs[lbi].AddItems(4, 5, 6)

		if lbs[lbi].Items[5] != 6 {
			t.Errorf("expected %d, got %d", 6, lbs[lbi].Items[5])
		}
	}
}

func TestLoadbalancer_Reset(t *testing.T) {
	set := []int{1, 2, 3}

	lbs := []Loadbalancer[int]{*NewLoadbalancer(set), *NewThreadSafeLoadbalancer(set)}

	for lbi := range lbs {

		for i := 0; i < 10; i++ {
			if lbs[lbi].Next() != set[i%len(set)] {
				t.Errorf("expected %d, got %d", set[i%len(set)], lbs[lbi].Next())
			}
		}

		lbs[lbi].Reset()

		if lbs[lbi].idx != 0 {
			t.Errorf("expected %d, got %d", 0, lbs[lbi].idx)
		}
	}
}

func TestLoadbalancer_Current(t *testing.T) {
	set := []int{1, 2, 3}

	lbs := []Loadbalancer[int]{*NewLoadbalancer(set), *NewThreadSafeLoadbalancer(set)}

	for lbi := range lbs {
		if lbs[lbi].Current() != set[0] {
			t.Errorf("expected %d, got %d", set[0], lbs[lbi].Current())
		}
	}
}
