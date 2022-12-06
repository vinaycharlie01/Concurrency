package main

import "fmt"

func genMes(ch1 chan string) {
	ch1 <- "Hello world"
}

func RealyMsg(ch chan string, ch2 chan string) {
	m := <-ch
	ch2 <- m

}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMes(ch1)
	go RealyMsg(ch1, ch2)

	v := <-ch2
	fmt.Println(v)
}
