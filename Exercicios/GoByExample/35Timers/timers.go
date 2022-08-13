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

	// time.Sleep(2 * time.Second)
}
