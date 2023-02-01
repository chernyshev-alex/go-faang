package dave

import "fmt"

func do_ch1() {
	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
