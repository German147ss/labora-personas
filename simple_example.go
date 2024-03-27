package main

import (
	"fmt"
	"sync"
)

func main10() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		mu.Lock()
		fmt.Println("Goroutine 1")
		wg.Done()
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		fmt.Println("Goroutine 2")
		wg.Done()
		mu.Unlock()
	}()

	wg.Wait()
}
