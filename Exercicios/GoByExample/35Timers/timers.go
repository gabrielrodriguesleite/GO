package main

import (
	"fmt"
	"time"
)

// As vezes precisamos executar um código depois de um período ou
// repetidamente com o mesmo intervalo. Go possui na biblioteca time
// as funções timer e ticker, ambas feitas para tornar essas tarefas mais fáceis.

func main() {

	// Timers representam um evento único no futuro.
	// Indicamos quanto tempo deve aguradar e ele entrega um canal que será
	// notificado quando completar o tempo. Este timer aguarda 2 segundos.
	timer1 := time.NewTimer(2 * time.Second)

	// O <-timer1.C bloqueia a execução até que a notificação do timer seja lançada.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Para o caso de simplesmente esperar é possível usar o time.Sleep.
	// Uma das razões do timer ser tão útil é a capacidade de cancelar
	// o timer antes dele disparar.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Dar tempo suficiente para o timer2 iniciar, se é que vai disparar, para
	// mostrar que está de fato parado.
	time.Sleep(2 * time.Second)
}

// O primeiro timer vai iniciar ~2s após o início do programa, mas o segundo
// deve ser parado antes de ter a chance de iniciar.
