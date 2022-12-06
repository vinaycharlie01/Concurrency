package main

import (
	"fmt"
	"sync"
)

func main() {
	// a := func(a int) int {
	// 	var i int
	// 	b := func(a int) int {
	// 		i++
	// 		return i
	// 	}
	// 	//return b(a)
	// 	fmt.Println("return from function")
	// 	return b(a)
	// }
	// fmt.Println(a(1))
	//a(1)
	//fmt.Println(a(1))

	//excisie 2
	// var wg sync.WaitGroup
	// //wg.Add(1)
	// incr := func(wg *sync.WaitGroup) {
	// 	var i int
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		i++
	// 		fmt.Printf("value of i : %v\n", i)
	// 	}()
	// 	fmt.Println("return from function")
	// 	return
	// }
	// incr(&wg)
	// wg.Wait()
	// fmt.Println("Done...!")

	//excisie 3

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	fmt.Println(i)

		// }()

		go func(i int) {
			defer wg.Done()
			fmt.Println(i)

		}(i)
	}
	wg.Wait()
}
