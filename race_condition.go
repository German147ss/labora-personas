// go run -race race_condition.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	var saldo int = 17

	wg.Add(2)
	go func() { // person a
		ac := saldo
		time.Sleep(1 * time.Nanosecond) // imagenos que se hace alguna operación larga
		ac = ac + 1
		saldo = ac
		wg.Done()
	}()
	go func() { // person b
		ac := saldo
		time.Sleep(1 * time.Nanosecond) // imagenos que se hace alguna operación larga
		ac = ac + 1
		saldo = ac
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("saldo final", saldo)
}
