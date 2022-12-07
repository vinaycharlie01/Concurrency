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
