package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só depois de lido")
}

func main() {
	c := make(chan int) // sem buffer

	go rotina(c)
	fmt.Println(<-c)
	fmt.Println("Lido")
	fmt.Println(<-c)
	fmt.Println("Deadlock..")
}
