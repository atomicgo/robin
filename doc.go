/*
Package robin is a simple, generic round-robin load balancer for Go.

It can be used to load balance any type of data. It is not limited to HTTP requests.

Robin takes any slice as an input and returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.

There are two versions of Robin: a thread-safe version (NewThreadSafeLoadbalancer)  and a non-thread-safe (NewLoadbalancer) version.
The thread-safe version is slower than the non-thread-safe version, but it is guaranteed that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

Benchmark:

	BenchmarkLoadbalancer_Next-32                   222961711                5.272 ns/op
	BenchmarkLoadbalancer_Next_ThreadSafe-32        79443891                15.57 ns/op
*/
package robin
