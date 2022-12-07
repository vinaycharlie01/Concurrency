---
# Golang Concurrency
---

### Concurrency 
+ Concurrency is about multiple things happening at same time in random order
+ Concurrent programs may or may not run in parallel. Concurrency is more about the structure of a program, which enables it to potentially run in parallel.
+ Go provides a built-in support for concurrency.
#### ex
+ Every morning I drink my coffee while I fix my breakfast. I could say I am working on both tasks concurrently, but I’m not actually executing them simultaneously (in parallel). At any given point in time I am either drinking from my cup or cutting pancakes, but never both at the same moment. I’m bouncing back-and-forth between the two activities concurrently.

### Parallelim
+ Parallelism is when multiple processes are carried out simultaneously by multiple threads or processors.
#### ex
+ Every morning I drink my coffee while watching my Tv. I could say I am working on both  in Parallel(im dring also and watching tv also).

---

+ Concurrency involves structuring a program so that `two or more tasks may be in progress simultaneously`, whereas parallelism allows for `two or more tasks to be executed simultaneously`. `SNote that while parallelism requires more than one processor or thread, concurrency does not`.

### Go’s Concurrency Tool Set
1. goroutines
2. channels
3. select
4. sync package


#### Normal Functions
```go

package main

import (
	"fmt"
)

func DringkingCoffee(a string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Dringking coffee,index %d and name %s\n", i, a)
	}
}

func Watching_Tv(a string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Waching tv ,index %d and name %s\n", i, a)

	}
}

func main() {
	var a string = "Vinay"
	DringkingCoffee(a)
	Watching_Tv(a)

}
```
```bash
output:
Dringking coffee,index 0 and name Vinay
Dringking coffee,index 1 and name Vinay
Dringking coffee,index 2 and name Vinay
Waching tv ,index 0 and name Vinay
Waching tv ,index 1 and name Vinay
Waching tv ,index 2 and name Vinay
```

we ne concurrently watching tv and  drinking coffee may or may not be same time but we need working  concurrently
we can achieve  using goroutines in golang

### Goroutine
+ Goroutine works like the thread in java
+ We can think Goroutines as user space threads 
`managed by go runtime`.
+ Goroutines extremely lightweight. Goroutines starts with 
2KB of stack, which grows and shrinks as required.
+ Can create hundreds of thousands of goroutines in the 
same address space.
+ Channels are used for communication of data between 
goroutines. Sharing of memory can be avoided.

syntax
```go
   go function name
   go fun1()
   go fun2()
   go fun3()
```


### What is a Thread?
+ Threads are smallest unit of execution that CPU accepts.
+ Process has atleast one thread - main thread.
+ Process can have multiple threads.
+ Threads share same address space.


### What are advantages of goroutines over OS threads?
+ Goroutine are extremely lightweight compared to OS 
threads.
+ Stack size is very small of 2kb as opposed to 8MB of stack 
of OS threads.
+ Context switching is very cheap as it happens in user 
space, goroutines have very less state to be stored.
+ Houndreds of thousands of goroutines can be created on 
single machine.

#### ex
```go
package main

import (
	"fmt"
	"time"
)

func DringkingCoffee(a string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Dringking coffee,index %d and name %s\n", i, a)
		time.Sleep(1 * time.Millisecond)
	}
}

func Watching_Tv(a string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Waching tv ,index %d and name %s\n", i, a)
		time.Sleep(1 * time.Millisecond)

	}
}

func main() {
	var a string = "Vinay"
	go DringkingCoffee(a)
	go Watching_Tv(a)
	time.Sleep(2 * time.Second)
}
```
```go
$ go run main.go
Waching tv ,index 0 and name Vinay
Dringking coffee,index 0 and name Vinay
Dringking coffee,index 1 and name Vinay
Waching tv ,index 1 and name Vinay
Waching tv ,index 2 and name Vinay
Dringking coffee,index 2 and name Vinay
```

### we can do better way by using `sync pkgs  waitgroup func`
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func DringkingCoffee(a string, wg *sync.WaitGroup) {
	//after completing loop we need to tell the waiting its done
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("Dringking coffee,index %d and name %s\n", i, a)
		time.Sleep(1 * time.Millisecond)
	}
}

func Watching_Tv(a string, wg *sync.WaitGroup) {
	//wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("Waching tv ,index %d and name %s\n", i, a)
		time.Sleep(1 * time.Millisecond)

	}
	wg.Done()
}

func main() {
	var a string = "Vinay"
	var wg sync.WaitGroup

	//How many go routines running
	wg.Add(2)
	go DringkingCoffee(a, &wg)
	go Watching_Tv(a, &wg)
	//time.Sleep(2 * time.Second)
	//then main() function waiting for go routines function  we need to wait for go routines function joined to the main function
	wg.Wait()

}
```
```go
$ go run main.go
Waching tv ,index 0 and name Vinay
Dringking coffee,index 0 and name Vinay
Dringking coffee,index 1 and name Vinay
Waching tv ,index 1 and name Vinay
Waching tv ,index 2 and name Vinay
Dringking coffee,index 2 and name Vinay
```