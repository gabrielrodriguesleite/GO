package main

import "fmt"

// Por padrão channels não são bufferizados, o que significa
// que eles podem apenas aceitar envios (chan <-) se eles
// tiverem um receptor (<- chan) pronto para receber o valor
// enviado. Buffered Channels aceitam um número limitado de valores
// sem ter um receptor correspondente para esses valores.
func main() {

	// Aqui criamos um channel que pode lidar com até 2 valores por vez
	messages := make(chan string, 2)

	// Por este channel ser buffered podemos enviar estes valores no channel
	// sem necessariamente termos um correspondente receptor concorrente.
	messages <- "buffered"
	messages <- "channel"

	// Em seguida podemos receber estes 2 valores normalmente
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
