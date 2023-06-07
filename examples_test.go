package robin_test

import (
	"atomicgo.dev/robin"
	"fmt"
)

func ExampleNewLoadbalancer() {
	set := []string{"object1", "object2", "object3"}
	lb := robin.NewLoadbalancer(set)

	fmt.Println(lb.Current())

	// Output:
	// object1
}

func ExampleLoadbalancer_Current() {
	set := []int{1, 2, 3}
	lb := robin.NewLoadbalancer(set)

	fmt.Println(lb.Current())

	// Output:
	// 1
}

func ExampleLoadbalancer_AddItems() {
	set := []int{1, 2, 3}
	lb := robin.NewLoadbalancer(set)

	lb.AddItems(4, 5, 6)

	fmt.Println(lb.Items)

	// Output:
	// [1 2 3 4 5 6]
}

func ExampleLoadbalancer_Reset() {
	set := []int{1, 2, 3, 4, 5, 6}
	lb := robin.NewLoadbalancer(set)

	lb.Next()
	lb.Next()
	lb.Next()

	lb.Reset()

	fmt.Println(lb.Current())

	// Output:
	// 1
}

func ExampleLoadbalancer_Next() {
	set := []int{1, 2, 3}
	lb := robin.NewLoadbalancer(set)

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
