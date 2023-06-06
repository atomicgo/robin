<h1 align="center">AtomicGo | robin</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Frobin&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/robin/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/robin?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/robin" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/robin/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/robin" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/robin?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/robin">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-5-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/robin" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/robin?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/robin#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/robin</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# robin

```go
import "atomicgo.dev/robin"
```

Package robin is a simple, generic round\-robin load balancer for Go.

It can be used to load balance any type of data. It is not limited to HTTP requests.

Robin takes any slice as an input and returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.

There are two versions of Robin: a thread\-safe version \(NewThreadSafeLoadbalancer\)  and a non\-thread\-safe \(NewLoadbalancer\) version. The thread\-safe version is slower than the non\-thread\-safe version, but it is guaranteed that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

Benchmark:

```
BenchmarkLoadbalancer_Next                      225866620                5.274 ns/op
BenchmarkLoadbalancer_Next-2                    227712583                5.285 ns/op
BenchmarkLoadbalancer_Next-32                   228792201                5.273 ns/op
BenchmarkLoadbalancer_Next_ThreadSafe           100000000               10.15 ns/op
BenchmarkLoadbalancer_Next_ThreadSafe-2         100000000               10.02 ns/op
BenchmarkLoadbalancer_Next_ThreadSafe-32        100000000               10.06 ns/op
```

## Index

- [type Loadbalancer](<#type-loadbalancer>)
  - [func NewLoadbalancer[T any](items []T) *Loadbalancer[T]](<#func-newloadbalancer>)
  - [func NewThreadSafeLoadbalancer[T any](items []T) *Loadbalancer[T]](<#func-newthreadsafeloadbalancer>)
  - [func (l *Loadbalancer[T]) AddItems(items ...T)](<#func-loadbalancert-additems>)
  - [func (l *Loadbalancer[T]) Next() T](<#func-loadbalancert-next>)
  - [func (l *Loadbalancer[T]) Reset()](<#func-loadbalancert-reset>)


## type [Loadbalancer](<https://github.com/atomicgo/robin/blob/main/robin.go#L6-L12>)

Loadbalancer is a simple, generic round\-robin load balancer for Go.

```go
type Loadbalancer[T any] struct {
    Items        []T
    CurrentIndex int
    ThreadSafe   bool
    // contains filtered or unexported fields
}
```

### func [NewLoadbalancer](<https://github.com/atomicgo/robin/blob/main/robin.go#L17>)

```go
func NewLoadbalancer[T any](items []T) *Loadbalancer[T]
```

NewLoadbalancer creates a new Loadbalancer. For maximum speed, this is not thread\-safe. Use NewThreadSafeLoadbalancer if you need thread\-safety. If two goroutines call Loadbalancer.Next at the exact same time, it can happen that they both return the same item.

### func [NewThreadSafeLoadbalancer](<https://github.com/atomicgo/robin/blob/main/robin.go#L26>)

```go
func NewThreadSafeLoadbalancer[T any](items []T) *Loadbalancer[T]
```

NewThreadSafeLoadbalancer creates a new Loadbalancer. This is thread\-safe, but slower than NewLoadbalancer. It is guaranteed that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

### func \(\*Loadbalancer\[T\]\) [AddItems](<https://github.com/atomicgo/robin/blob/main/robin.go#L59>)

```go
func (l *Loadbalancer[T]) AddItems(items ...T)
```

AddItems adds items to the Loadbalancer.

### func \(\*Loadbalancer\[T\]\) [Next](<https://github.com/atomicgo/robin/blob/main/robin.go#L35>)

```go
func (l *Loadbalancer[T]) Next() T
```

<details><summary>Example</summary>
<p>

```go
{
	set := []int{1, 2, 3}
	lb := NewLoadbalancer(set)

	for i := 0; i < 10; i++ {
		fmt.Println(lb.Next())
	}

}
```

#### Output

```
1
2
3
1
2
3
1
2
3
1
```

</p>
</details>

### func \(\*Loadbalancer\[T\]\) [Reset](<https://github.com/atomicgo/robin/blob/main/robin.go#L50>)

```go
func (l *Loadbalancer[T]) Reset()
```

Reset resets the Loadbalancer to its initial state.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
