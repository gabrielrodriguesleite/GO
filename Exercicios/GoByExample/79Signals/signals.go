package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// As vezes precisamos que nossa aplicação reaja aos sinais do Unix.
// Mais em: https://en.wikipedia.org/wiki/Unix_signal
// Por exemplo, quando temos um serviço rodando e queremos que ele finalize
// corretamente quando recebe um SIGTERM, ou uma ferramenta de linha de comando
// que pare de processar a entrada ao receber um SIGINT.
// Neste exemplo veremos como manejar sinais do sistema com canais Go.
func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println("\nSinal:", sig)
		done <- true
	}()

	fmt.Println("Aguardando sinal...")
	<-done
	fmt.Println("Sinal recebido\nSaindo.")
}
