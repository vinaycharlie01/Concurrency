package main

import (
	"fmt"
	"time"
)

//Go Routines as user space threads managed by go runtime
//Go: Concurrency vs Parallelism - Medium

func Fun(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	//Direct Call
	//Fun("Direct Call")

	//go routines fun call
	//go Fun("Go Routines")

	//go routines with anonymous func
	go func() {
		Fun("anonymous func call")
	}()
	go func() {
		Fun("anonymous func call2")
	}()
	fmt.Println("waiting for Go Routines")

	// using  func value call
	a := Fun
	go a("func value call")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Done")

}
