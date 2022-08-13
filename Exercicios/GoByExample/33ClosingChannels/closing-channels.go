package main

import "fmt"

// Fechar (close(channel)) um channel indica que
// não serão enviados novos valores por ele
// Isso pode ser util para comunicar para o receptor
// que a transmissão está completa.

func main() {

	// Neste exemplo é usado 2 channels, job comunica o serviço
	// feito a partir da main() goroutine para a worker goroutine.
	// Quando não temos mais jobs para o worker nos fechamos o
	// channel jobs com close(jobs)
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Este goroutine worker recebe repetidas vezes de jobs com
	// j, more := <-jobs
	// Nesta configuração o valor de more vai ser false se jobs
	// tiver sido fechado e todos os valores no channel tiverem
	// sido recebidos. Usamos more para notificar quando o
	// worker (goroutine) recebeu todos os nossos jobs
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Este for envia 3 jobs para o worker pelo channel jobs então o fecha
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	// Usando o padrão visto anteriormente de sincronização
	// para aguardarmos o worker completar
	<-done

}

/*
 	SAÍDA DO EXEMPLO MAS A PARTIR DA V 1.18 SAI DIFERENTE INCLUSIVE NO PLAYGROUND
$ go run closing-channels.go
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
*/
