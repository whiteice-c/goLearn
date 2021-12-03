package channel

import (
	"fmt"
	"time"
)

func producer(max int, prod chan<- int) {
	for {
		for i := 0; i < max; i++ {
			prod <- i
			fmt.Println("producer:", i)
			time.Sleep(1 * time.Second)
		}
	}
}

func consumer(cons <-chan int, exitSig chan<- int) {
	for {
		v := <-cons
		fmt.Println("consumer:", v)
		time.Sleep(2 * time.Second)
	}

	exitSig <- 1
}

func TestModel() {
	prod := make(chan int, 10)
	cons := make(chan int, 10)
	exitSig := make(chan int)

	go producer(20, prod)
	go consumer(cons, exitSig)

	<-exitSig
	fmt.Println("main exit")
}
