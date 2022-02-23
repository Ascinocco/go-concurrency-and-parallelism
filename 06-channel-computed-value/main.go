package main

import "fmt"

func chanAdd() {
	ch := make(chan int)

	go func(a, b int) {
		ch <- a + b
	}(1, 2)

	fmt.Println("chan result", <-ch)
}

func chanRange() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}

		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func buffChan() {
	ch := make(chan int, 3)

	go func() {
		defer close(ch)

		for i := 0; i < 6; i++ {
			fmt.Println("Sent", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println("Received", v)
	}
}

func genMsg(ch chan<- string) {
	ch <- "hi"
}
func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	m := <-ch1
	ch2 <- m
}

func chanDir() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMsg(ch1)

	go relayMsg(ch1, ch2)

	fmt.Println("Message relayed:", <-ch2)
}

func chanOwner() {
	owner := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)

			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Println("Received:", v)
		}

		fmt.Println("Done receiving")
	}

	ch := owner()
	consumer(ch)
}

func main() {
	chanOwner()
}
