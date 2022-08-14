package main

import (
	"fmt"
	"time"
)

// Aqui temos um exemplo de como implementar um worker pool usando
// goroutines e channels

// Aqui um worker do qual iremos rodar varias instancias concorrentes.
// Estes workers vão receber trabalhos pelo canal jobs e enviar
// o resultado pelo canal results.
// Vamos usar sleep por 1s por trabalho para simular uma tarefa pesada
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
