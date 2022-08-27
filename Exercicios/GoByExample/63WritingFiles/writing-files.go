package main

import (
	"fmt"
	"os"
)

// Escrever arquivo usando Go é uma tarefa parecida com a de leitura
// demonstrada no exemplo anterior

// Novamente é muito útil ter uma função para tratar a verificação por erros
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Exemplo de escrita de string ou bytes em um arquivo.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

}
