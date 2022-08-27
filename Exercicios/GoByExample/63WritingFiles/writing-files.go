package main

import (
	"bufio"
	"fmt"
	"os"
)

// Escrever arquivo usando Go é uma tarefa parecida com a de leitura
// demonstrada no exemplo anterior.

// Novamente é muito útil ter uma função para tratar a verificação por erros.
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

	// Para escritas mais granulares, primeiro abrimos o arquivo para escrita:
	f, err := os.Create("/tmp/dat2")
	check(err)

	// O padrão da linguagem é agendar o fechamento logo após abri o arquivo.
	defer f.Close()

	// É possível escrever slices de bytes como esperado.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString permite escrever uma string diretamente.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Chamar Sync escreve os dados no armazenamento rom liberando memória ram.
	f.Sync()

	// bufio fornece escrita buferizada.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes \n", n4)

	w.Flush()
}
