package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// No exemplo anterior foi usado mutex para bloquear o acesso explicitamente
// e assim sincronizar o acesso ao estado compartilhado entre muitas
// goroutines. Ainda outra opção é usar o recurso de sincronização de
// goroutines e channels para conseguir o mesmo resultado.
// Esta abordagem baseada em canais se alinha com a idéia de Go de compartilhar
// memoria por comunicar e ter uma parte dos dados pertencente a apenas uma
// goroutine.

// Neste exemplo o estrado vai pertencer a apenas uma goroutine.
// Isto vai garantir que o dado nunca seja corrompido por um acesso concorrente.
// Para ter acesso de leitura e escrita, as goroutines enviarão mensagens para
// a goroutine que possui o dado e esta enviará as correspondentes respostas.
// As estruturas readOp e writeOp encapsulam estas vias de requisições por onde
// a goroutine proprietária possa responder.
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Vamos registrar quantos operações serão realizadas.
	var readOps uint64
	var writeOps uint64

	// Os canais reads e writes vão ser usados por outras goroutines para lançar
	// as requisições de leitura e escrita.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// A nossas goroutine que possui o state, que será um map assim como no exemplo
	// anterior, porém agora privado para a goroutine stateful.
	// Esta goroutine usa um select em loop entre os canais reads e writes,
	// respondendo as requisições que chegam.
	// Uma resposta é executada primeiramente performando uma operação de requisição
	// e então enviando um valor no canal de resposta resp para indicar sucesso
	// (e o valor desejado no caso de leitura reads)
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Aqui é iniciado 100 goroutines que lançaram requisições de leitura para a
	// goroutine proprietária do estado via canal reads.
	// Cada read precisa construir uma estrutura readOp e enviar através do canal
	// reads e então receber o resultado pelo canal de resposta provido, resp.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Iniciamos também 10 goroutines de escrita usando uma abordagem similar.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Permitir que as goroutines trabalhem por 1 segundo.
	time.Sleep(time.Second)

	// Capturar as contagens e reportas os resultados.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

// Rodando a aplicação mostra que esta abordagem de manejamento de estado
// por meio de uma goroutine dona do estado completou mais de 80.000 operações
// no total em 1 segundo.
