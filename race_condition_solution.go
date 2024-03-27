package main

import (
	"fmt"
	"sync"
)

func main2() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Inicializar la WaitGroup con 2 goroutines
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	fmt.Println("Contador final:", counter)
}
