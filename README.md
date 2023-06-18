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
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-11-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
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

Package robin is a simple, generic and thread\-safe round\-robin load balancer for Go.

It can be used to load balance any type of data. It is not limited to HTTP requests.

Robin takes any slice as an input and returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.

Thread\-safety is achieved by using atomic operations amd guarantees that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

Benchmark:

```
BenchmarkLoadbalancer_Next              252151534                4.746 ns/op           0 B/op          0 allocs/op
BenchmarkLoadbalancer_Next-2            254281032                4.758 ns/op           0 B/op          0 allocs/op
BenchmarkLoadbalancer_Next-4            253424396                4.738 ns/op           0 B/op          0 allocs/op
BenchmarkLoadbalancer_Next-8            254842484                4.752 ns/op           0 B/op          0 allocs/op
BenchmarkLoadbalancer_Next-16           247016046                4.785 ns/op           0 B/op          0 allocs/op
BenchmarkLoadbalancer_Next-32           250539441                4.774 ns/op           0 B/op          0 allocs/op
```

## Index

- [type Loadbalancer](<#Loadbalancer>)
  - [func NewLoadbalancer\[T any\]\(items \[\]T\) \*Loadbalancer\[T\]](<#NewLoadbalancer>)
  - [func \(l \*Loadbalancer\[T\]\) AddItems\(items ...T\)](<#Loadbalancer[T].AddItems>)
  - [func \(l \*Loadbalancer\[T\]\) Current\(\) T](<#Loadbalancer[T].Current>)
  - [func \(l \*Loadbalancer\[T\]\) Next\(\) T](<#Loadbalancer[T].Next>)
  - [func \(l \*Loadbalancer\[T\]\) Reset\(\)](<#Loadbalancer[T].Reset>)


<a name="Loadbalancer"></a>
## type [Loadbalancer](<https://github.com/atomicgo/robin/blob/main/robin.go#L8-L12>)

Loadbalancer is a simple, generic and thread\-safe round\-robin load balancer for Go.

```go
type Loadbalancer[T any] struct {
    Items []T
    // contains filtered or unexported fields
}
```

<details><summary>Example (Demo)</summary>
<p>



```go
package main

import (
	"atomicgo.dev/robin"
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	people := []Person{
		{Name: "person1"},
		{Name: "person2"},
		{Name: "person3"},
		{Name: "person4"},
		{Name: "person5"},
	}

	lb := robin.NewLoadbalancer(people)

	for i := 0; i < 10; i++ {
		fmt.Println(lb.Next().Name)
	}

}
```

#### Output

```
person1
person2
person3
person4
person5
person1
person2
person3
person4
person5
```

</p>
</details>

<a name="NewLoadbalancer"></a>
### func [NewLoadbalancer](<https://github.com/atomicgo/robin/blob/main/robin.go#L16>)

```go
func NewLoadbalancer[T any](items []T) *Loadbalancer[T]
```

NewLoadbalancer creates a new Loadbalancer. It is guaranteed that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/robin"
	"fmt"
)

func main() {
	set := []string{"object1", "object2", "object3"}
	lb := robin.NewLoadbalancer(set)

	fmt.Println(lb.Current())

}
```

#### Output

```
object1
```

</p>
</details>

<a name="Loadbalancer[T].AddItems"></a>
### func \(\*Loadbalancer\[T\]\) [AddItems](<https://github.com/atomicgo/robin/blob/main/robin.go#L40>)

```go
func (l *Loadbalancer[T]) AddItems(items ...T)
```

AddItems adds items to the Loadbalancer.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/robin"
	"fmt"
)

func main() {
	set := []int{1, 2, 3}
	lb := robin.NewLoadbalancer(set)

	lb.AddItems(4, 5, 6)

	fmt.Println(lb.Items)

}
```

#### Output

```
[1 2 3 4 5 6]
```

</p>
</details>

<a name="Loadbalancer[T].Current"></a>
### func \(\*Loadbalancer\[T\]\) [Current](<https://github.com/atomicgo/robin/blob/main/robin.go#L23>)

```go
func (l *Loadbalancer[T]) Current() T
```

Current returns the current item in the slice, without advancing the Loadbalancer.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/robin"
	"fmt"
)

func main() {
	set := []int{1, 2, 3}
	lb := robin.NewLoadbalancer(set)

	fmt.Println(lb.Current())

}
```

#### Output

```
1
```

</p>
</details>

<a name="Loadbalancer[T].Next"></a>
### func \(\*Loadbalancer\[T\]\) [Next](<https://github.com/atomicgo/robin/blob/main/robin.go#L29>)

```go
func (l *Loadbalancer[T]) Next() T
```

Next returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/robin"
	"fmt"
)

func main() {
	set := []int{1, 2, 3}
	lb := robin.NewLoadbalancer(set)

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

<a name="Loadbalancer[T].Reset"></a>
### func \(\*Loadbalancer\[T\]\) [Reset](<https://github.com/atomicgo/robin/blob/main/robin.go#L35>)

```go
func (l *Loadbalancer[T]) Reset()
```

Reset resets the Loadbalancer to its initial state.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/robin"
	"fmt"
)

func main() {
	set := []int{1, 2, 3, 4, 5, 6}
	lb := robin.NewLoadbalancer(set)

	lb.Next()
	lb.Next()
	lb.Next()

	lb.Reset()

	fmt.Println(lb.Current())

}
```

#### Output

```
1
```

</p>
</details>

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
