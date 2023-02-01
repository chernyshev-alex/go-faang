package dave

import (
	"fmt"
	"time"
)

func do_ch3() {
	a, b := make(chan bool), make(chan bool)
	t0 := time.Now()
	go func() {
		close(a)
		close(b)
	}()
	WaitMany(a, b)
	fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
}

func WaitMany(a chan bool, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil

		}
	}
}
