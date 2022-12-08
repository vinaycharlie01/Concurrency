// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Learnig(a string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("learnig", a)
// 	//time.Sleep(1 * time.Millisecond)
// }

// func UsingMobile(s string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("UsingMobile", s)
// 	//time.Sleep(1 * time.Millisecond)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		//go Learnig("Vinay",&wg)
// 		go func() {
// 			Learnig("Vinay", &wg)
// 		}()
// 	}
// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		go func() {
// 			UsingMobile("Vinay", &wg)
// 		}()

// 	}
// 	wg.Wait()

// }

//2
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Add(a chan int, wg *sync.WaitGroup) {
// 	//defer close(a)
// 	// defer close(a)
// 	defer close(a)
// 	defer wg.Done()
// 	//defer close(a)
// 	a <- 10
// }

// func sub(b chan int, wg *sync.WaitGroup) {

// 	defer wg.Done()
// 	//defer close(b)

// 	b <- 20

// }

// func main() {
// 	ch := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go Add(ch, &wg)
// 	go sub(ch, &wg)
// 	//wg.Wait()
// 	for i := 0; i < 2; i++ {
// 		fmt.Println(<-ch)
// 	}
// 	wg.Wait()
// }

// //3
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // func Credited(a chan int, wg *sync.WaitGroup) {
// // 	defer wg.Done()

// // }

// func main() {
// 	ch := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		ch <- 1
// 		ch <- +2
// 		defer close(ch)
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case m := <-ch:
// 			fmt.Println(m)
// 		case m := <-ch:
// 			fmt.Println(m)

// 		}
// 	}
// 	wg.Wait()

// }

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		////defer wg.Done()
		ch <- 1
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	// wg.Wait()
	fmt.Println("Hello world!")
}
