package main

import "fmt"

// Por padrão enviar e receber por channels são operações bloqueantes.
// Contudo, por usar um select com caso default implementamos
// envios e recepções não bloqueantes, mesmo para select
// multi direcionais sem bloqueios.

func main() {

	messages := make(chan string)
	// signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

}
