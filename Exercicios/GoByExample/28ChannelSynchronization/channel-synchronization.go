package main

import (
	"fmt"
	"time"
)

// Podemos usar channels para sincronizar a execução entre goroutines
// Neste exemplo usamos um receptor bloqueante para aguardarmos até
// uma goroutine terminar.
// (Para aguardar várias goroutines é preferível usar WaitGroup)

// Esta função será executada em uma goroutine ela leva > 1s para finalizar.
// O channel done será usado para notificar outra goroutine que esta função
// finalizou.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// enviando um valor para notificar que finalizou
	done <- true
}

func main() {

	// Iniciando a goroutine worker, passando um channel usado para notificar
	done := make(chan bool, 1)
	go worker(done)

	// Execução bloqueada até que o transmissor envie o valor.
	// Execução aguarda até receber algum valor.
	// Removendo <-done do código o programa irá sair antes mesmo do worker iniciar
	<-done
}
