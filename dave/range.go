package dave

import "fmt"

func do_range() {
	q := make(chan string, 2)
	q <- "1"
	q <- "2"
	close(q)

	for e := range q {
		fmt.Printf("%s\n", e)
	}
}
