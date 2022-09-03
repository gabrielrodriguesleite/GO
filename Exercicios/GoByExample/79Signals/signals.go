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

	// Notificações de sinais Go funcionam por enviar valores "os.Signal" em um canal.
	// Um canal é criado para receber estas notificações.
	// Note que este canal deve ser bufferizado.
	sigs := make(chan os.Signal, 1)

	// "signal.Notify" registra o canal fornecido para receber nitificações para
	// os sinais especificados.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// É possível receber de "sigs" aqui na função principal, mas neste exemplo
	// será feito em uma goroutine para demonstrar um cenário mais realístico
	// de desligamento elegante.
	done := make(chan bool, 1)

	go func() {
		// A goroutine executa um bloco esperando receber sinais.
		// Quando receber um irá imprimi-lo e então irá notificar o programa
		// que este pode terminar.
		sig := <-sigs
		fmt.Println("\nSinal:", sig)
		done <- true
	}()

	fmt.Println("Aguardando sinal...")
	<-done
	fmt.Println("Sinal recebido\nSaindo.")
}

// SAÍDA ESPERADA PARA ESTE CÓDIGO:

/* go run 79Signals/signals.go
Aguardando sinal...
^C
Sinal: interrupt
Sinal recebido
Saindo.
*/
