package main

import (
	"fmt"
	"sync"
)

func main4() {
	// Creamos dos mutexes
	mu1 := &sync.Mutex{}
	mu2 := &sync.Mutex{}

	// Goroutine 1
	go func() {
		// Goroutine 1 adquiere el mutex 1
		mu1.Lock()
		fmt.Println("Goroutine 1 adquirió el mutex 1")

		// Goroutine 1 intenta adquirir el mutex 2
		mu2.Lock()
		fmt.Println("Goroutine 1 adquirió el mutex 2")

		// Libera los mutexes en orden inverso
		mu2.Unlock()
		mu1.Unlock()
	}()

	// Goroutine 2
	go func() {
		// Goroutine 2 adquiere el mutex 2
		mu2.Lock()
		fmt.Println("Goroutine 2 adquirió el mutex 2")

		// Goroutine 2 intenta adquirir el mutex 1
		mu1.Lock()
		fmt.Println("Goroutine 2 adquirió el mutex 1")

		// Libera los mutexes en orden inverso
		mu1.Unlock()
		mu2.Unlock()
	}()

	// El programa se quedará bloqueado aquí
	fmt.Scanln()
}
