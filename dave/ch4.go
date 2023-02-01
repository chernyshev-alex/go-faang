package dave

import (
	"fmt"
)

func do_ch4() {
	var c = make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	for v, ok := <-c; ok; v, ok = <-c {
		fmt.Printf("%d %v\n", v, ok)
	}
	// for v := range c {
	// 	fmt.Printf("%d ", v)
	// }
}
