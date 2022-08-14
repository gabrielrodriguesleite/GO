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

	// Para usar a pool de workers precisamos enviar trabalhos e coletar
	// resultados. Para isso usamos 2 canais, jobs e results.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Iniciamos 3 workers bloqueados pois se iniciam sem tarefas
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Então enviamos 5 serviços então fechamos o canal jobs indicando
	// que estes são todos os serviços que faremos.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finalmente coletamos todos os resultados dos trabalhos.
	// Também garantimos que as goroutines workers finalizaram.
	// Uma alternativa para aguardar multiplas goroutines é usando
	// WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}

}

// Esse código executado mostra os 5 jobs sendo executados pelos
// 3 workers. O programa leva em torno de 2 segundos ao invés de 5
// por conta de estar sendo executado paralelamente por 3 workers
