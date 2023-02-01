package dave

import (
	"fmt"
	"time"
)

func f(s string) {
	fmt.Println(s)
}

func do_ch10() {
	go f("a")
	go f("b")
	time.Sleep(time.Second)
}

func f1(c chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	c <- true
}

func do_ch10A() {
	pings, pongs := make(chan string, 1), make(chan string, 1)
	pings <- "p1"
	m1 := <-pings
	pongs <- m1
	fmt.Println(<-pongs)
}

func do_ch10B() {
	c1, c2 := make(chan string), make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "1"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		c2 <- "2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("got", m1)
		case m2 := <-c2:
			fmt.Println("got", m2)
		case <-time.After(2 * time.Second):
			fmt.Println("timeout 2")
		}
	}
}

// non blocking
func do_ch10C() {
	mq, signals := make(chan string), make(chan string)
	select {
	case msg := <-mq:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	select {
	case mq <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-mq:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func do_close() {
	jobs, done := make(chan int, 5), make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
