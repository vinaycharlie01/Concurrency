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
