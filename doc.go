/*
Package robin is a simple, generic and thread-safe round-robin load balancer for Go.

It can be used to load balance any type of data. It is not limited to HTTP requests.

Robin takes any slice as an input and returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.

Thread-safety is achieved by using atomic operations amd guarantees that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

Benchmark:

	BenchmarkLoadbalancer_Next              252151534                4.746 ns/op           0 B/op          0 allocs/op
	BenchmarkLoadbalancer_Next-2            254281032                4.758 ns/op           0 B/op          0 allocs/op
	BenchmarkLoadbalancer_Next-4            253424396                4.738 ns/op           0 B/op          0 allocs/op
	BenchmarkLoadbalancer_Next-8            254842484                4.752 ns/op           0 B/op          0 allocs/op
	BenchmarkLoadbalancer_Next-16           247016046                4.785 ns/op           0 B/op          0 allocs/op
	BenchmarkLoadbalancer_Next-32           250539441                4.774 ns/op           0 B/op          0 allocs/op
*/
package robin
