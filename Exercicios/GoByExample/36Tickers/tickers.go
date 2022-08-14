package main

import (
	"fmt"
	"time"
)

// Timers serve para quando queremos fazer algo uma vez no futuro.
// Ticker são para quando queremos repetir alguma ação dentro de intervalos regulares.
// Neste exemplo temos um tick que imprime um valor periodicamente até que o paremos.

func main() {

	// Ticker usam um mecanismo similare aos timers. Um channel que envia valores.
	// Nesse código usamos um select no channel do ticker para aguardar os
	// valores que chegam a cada ~500ms
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

	// Tickers podem ser parados assim como timers. Uma vez que um ticker é parado
	// ele não consegue mais receber nenhum valor no seu channel.
	// Aqui paramos após 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

// Quando rodamos este programa o ticker deve executar 3 vezes antes de ser parado.
