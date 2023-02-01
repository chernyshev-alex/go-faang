package dave

import (
	"fmt"
	"sync"
	"time"
)

func do_ch2() {
	finish := make(chan struct{})
	var done sync.WaitGroup
	for i := 0; i < 3; i++ {
		done.Add(1)
		go func(ng int) {
			fmt.Printf(" %d started\n", ng)
			select {
			case <-time.After(4 * time.Second):
			case <-finish:
			}
			done.Done()
		}(i)
	}
	t0 := time.Now()
	close(finish)
	done.Wait()
	fmt.Printf("Waited %v for goroutines to stop\n", time.Since(t0))
}
