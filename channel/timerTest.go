package channel

import (
	"fmt"
	"time"
)

func TimerTest() {
	ch := make(chan int, 10)
	timerPrd := time.NewTimer(1 * time.Second)
	timerCon := time.NewTimer(1 * time.Second)
	dat := 0

	for {
		select {
		case <-timerPrd.C:
			fmt.Printf("producer : %d\n", dat)
			ch <- dat
			dat++
			timerPrd.Reset(1 * time.Second)

		case <-timerCon.C:
			i := <-ch
			fmt.Printf("consumer : %d\n", i)
			timerCon.Reset(1 * time.Second)
			if i == 15 {
				return
			}
		}
	}

}
