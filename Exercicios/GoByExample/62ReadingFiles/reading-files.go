package main

import (
	"fmt"
	"io"
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

	// Para um controle melhor de quanto se deseja ler do arquivo ou de qual parte
	// começamos por abrir o arquivo para obter um valor os.File .
	f, err := os.Open("/tmp/dat")
	check(err)

	// Ler alguns bytes do inicio do arquivo.
	// Permitimos até 5 bytes serem lidos, e então verificamos quantos foram lidos de fato.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// É possível ler a partir de uma posição determinada no arquivo.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// A função ReadAtLeas permite ler como no exemplo anterior de uma forma mais robusta.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
}
