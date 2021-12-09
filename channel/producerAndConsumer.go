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
		time.Sleep(3 * time.Second)

		if v == 5 {
			fmt.Println("consumer exit")
			exitSig <- 1 //消费者消费到5后发出退出信号，通知主进程退出
			return
		}
	}
}

//TestModel 测试生产消费者模型
func TestModel() {
	fchan := make(chan int, 5)
	exitSig := make(chan int)

	go producer(10, fchan)
	go consumer(fchan, exitSig)

	<-exitSig
	fmt.Println("main exit")
}
