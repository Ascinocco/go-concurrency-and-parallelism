package main

import (
	"fmt"
	"time"
)

func printString(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	printString("direct call")

	go printString("go-routine-1")

	go func() {
		printString("go-routine-2")
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("done...")
}
