package main

import (
	"fmt"
	"sync"
)

func main5() {
	mu1 := &sync.Mutex{}
	mu2 := &sync.Mutex{}

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 adquirió el mutex 1")
		mu2.Lock()
		fmt.Println("Goroutine 1 adquirió el mutex 2")
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 adquirió el mutex 1")
		mu2.Lock()
		fmt.Println("Goroutine 2 adquirió el mutex 2")
		mu2.Unlock()
		mu1.Unlock()
	}()

	// El programa no se bloqueará
	fmt.Scanln()
}
