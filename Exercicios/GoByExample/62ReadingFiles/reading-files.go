package main

import (
	"fmt"
	"os"
)

// Ler e escrever arquivos são tarefas básicas porém necessárias para muitos
// programas Go. Primeiramente veremos exemplos de leitura.

// Ler arquivos envolve verificar por erros com frequência. Esta função
// irá ajudar nessa tarefa.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// A tarefa mais comum de leitura de um arquivo inteiro para a memória
	// e então imprimindo o conteúdo convertido para string na tela.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

}
