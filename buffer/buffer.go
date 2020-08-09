package main

import "fmt"

func rotina(ch chan int) {

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou")

	ch <- 5
	ch <- 6
	ch <- 7
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)
	fmt.Println(<-ch)
}
