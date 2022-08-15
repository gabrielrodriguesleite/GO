package main

import "fmt"

// É possível recuperar de um panic com a função recover.
// recover pode ser usado para evitar que a aplicação seja abortada
// por um panic e permitir que a aplicação continue rodando

// Um exemplo onde isso pode ser útil: um servidor não deve de
// quebrar se um cliente apresentar um erro crítico. Ao invés
// disso o servidor deve fechar a conexão e continuar servindo
// os outros clientes.
// De fato, isto é o que o servidor HTTP do pacote net/http faz.

// Esta função lança um panic
func mayPanic() {
	panic("a problem")
}

func main() {
	// recover deve ser chamado com uma função postergada com defer.
	// Quando a função main lança um panic, a função postergada vai
	// ativar uma chamada recover que irá capturar o panic
	// O valor retornado do recover é o erro levantado na chamada
	// de panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: \n", r)
		}
	}()

	mayPanic()

	// Esta próxima linha não será executada, já que mayPanic estoura um panic
	// Então a execução de main é parada no ponto onde o panic aconteceu
	// e segue para o ponto da cláusula defer
	fmt.Println("After mayPanic()")
}
