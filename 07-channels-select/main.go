package main

import (
	"fmt"
	"time"
)

func goSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)

		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}
}

func goSelectTimeout() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
		break
	}
}

func goSelectNonBlocking() {
	ch := make(chan string, 1)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "hi"
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no msg received")
		}
	}

	fmt.Println("processing")
	time.Sleep(1500 * time.Millisecond)
}

func main() {
	goSelectNonBlocking()
}
