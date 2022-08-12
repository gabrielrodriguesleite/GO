package main

// Uma goroutine é uma linha de execução leve
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suponha que temos uma função call `f(s)`
	// Aqui é como chamamos ela normalmente de maneira síncrona
	f("direct")

	// Para invocar numa goroutine chamamos com `go f(s)`
	// Esta nova goroutine vai executar de modo concorrente com a primeira chamada
	go f("goroutine")

	// É possível também invocar uma goroutine usando uma função anônima
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// As duas chamadas da função são executadas de forma assíncrona em
	// goroutines separadas. Aqui aguardamos para que terminem.
	// (Uma forma melhor seria usando WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")

}

// Quando esse código é executado vemos a execução da primeira chamada
// bloqueante primeiro, então a saída das duas goroutines.
// A saída das duas goroutines devem ser intercaladas por que as goroutines
// rodam de maneira concorrentes.
