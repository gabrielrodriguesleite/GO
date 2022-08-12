package main

import "fmt"

// Channels são tubos que conectam goroutines concorrentes.
// É possível enviar valores para channels de uma goroutine
// e receber estes valores em outra goroutines
func main() {

	// Cria-se um novo channel com make(chan val-type)
	// Channels são tipados pelos valores que eles comunicam
	messages := make(chan string)

	// Enviar um valor em um channel usando a sintaxe channel <- valor
	// Aqui enviamos o valor "pint" para o channel criado acima
	// a partir de uma nova goroutine
	go func() { messages <- "ping" }()

	// Para receber um valor vindo pelo channel samos a sintaxe <-channel
	// Aqui recebemos o valor enviado acima e imprimimos
	msg := <-messages
	fmt.Println(msg)
}

// Quando rodamos esse código "ping" é passado de uma goroutine
// para outra por meio de um channel

// Por padrão enviar e receber bloqueia até ambos quem envia e
// quem recebe estejam prontos. Esta propriedade permite aguardar
// até a finalização da aplicação para que a mensagem "ping" viage
// sem ter que usar outra forma de sincronização
