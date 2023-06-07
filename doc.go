/*
Package robin is a simple, generic and thread-safe round-robin load balancer for Go.

It can be used to load balance any type of data. It is not limited to HTTP requests.

Robin takes any slice as an input and returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.

Thread-safety is achieved by using atomic operations amd guarantees that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

Benchmark:

	BenchmarkLoadbalancer_Next              251751190                4.772 ns/op
	BenchmarkLoadbalancer_Next-2            250728889                4.834 ns/op
	BenchmarkLoadbalancer_Next-4            253328150                4.773 ns/op
	BenchmarkLoadbalancer_Next-8            248147372                4.783 ns/op
	BenchmarkLoadbalancer_Next-16           249468267                4.773 ns/op
	BenchmarkLoadbalancer_Next-32           247134729                4.802 ns/op
*/
package robin
