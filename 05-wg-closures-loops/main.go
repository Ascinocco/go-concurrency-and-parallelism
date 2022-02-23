package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(count int) {
			defer wg.Done()
			fmt.Println(count)
		}(i)
	}

	wg.Wait()
}
