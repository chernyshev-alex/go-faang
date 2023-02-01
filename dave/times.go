package dave

import (
	"fmt"
	"sync"
	"time"
)

func do_timer(wg *sync.WaitGroup) {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")
	timer1 = time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired again")
	wg.Done()
}

func do_ticker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()
	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
}
