package main

import (
	"fmt"
	"time"
)

// Timers serve para quando queremos fazer algo uma vez no futuro.
// Ticker são para quando queremos repetir alguma ação dentro de intervalos regulares.
// Neste exemplo temos um tick que imprime um valor periodicamente até que o paremos.

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
