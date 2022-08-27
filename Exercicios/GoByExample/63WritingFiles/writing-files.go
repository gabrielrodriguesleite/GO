package main

import "os"

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

}
