package main

import (
	"fmt"
	"time"
)

// As vezes precisamos executar um código depois de um período ou
// repetidamente com o mesmo intervalo. Go possui na biblioteca time
// as funções timer e ticker, ambas feitas para tornar essas tarefas mais fáceis.

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// O primeiro timer vai iniciar ~2s após o início do programa, mas o segundo
// deve ser parado antes de ter a chance de iniciar.
