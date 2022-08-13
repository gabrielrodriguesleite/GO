package main

import "fmt"

// Fechar (close(channel)) um channel indica que
// não serão enviados novos valores por ele
// Isso pode ser util para comunicar para o receptor
// que a transmissão está completa.

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

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

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

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
