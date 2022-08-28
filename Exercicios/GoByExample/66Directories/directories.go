package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Go possui varias funções úteis para trabalhar com pastas no sistema de arquivos
func main() {

	// Criando um subdiretório no diretório atual.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Quando se está trabalhando com diretórios temporários é uma boa prática
	// agendar sua remoção. os.RemoveAll vai deletar uma árvore de diretório
	// inteiro similar ao "rm -rf".
	defer os.RemoveAll("subdir")

	// Função auxiliar para criar arquivos vazios.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	// Criar um arquivo vazio.
	createEmptyFile("subdir/file1")

	// MkdirAll similar ao comando mkdir -p cria a hierarquia de diretórios.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lista o conteúdo do diretório, retornando uma
	// slice de objetos do tipo os.DirEntry
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent") // Listing subdir/parent
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
		// child true
		// file2 false
		// file3 false
	}

	// Chdir similar ao "cd", permite mudar o diretório de trabalho.
	err = os.Chdir("subdir/parent/child")

	check(err)

	// Agora para listar o conteúdo do diretório atual:
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child") // Listing subdir/parent/child
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
		// file4 false
	}

	// Navegando de volta para o início.
	err = os.Chdir("../../../")
	check(err)

	// É possível visitar um diretório recursivamente, incluindo todos os seus
	// subdiretórios. Walk aceita uma callback para lidar com cada arquivo ou
	// diretório visitado.
	fmt.Println("Visiting subdir") // Visiting subdir
	err = filepath.Walk("subdir", visit)

}

// A callback usada em Walk, chamada para cada arquivo ou diretório encontrado
// recursivamente pelo método Walk de filepath
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	// subdir true
	// subdir/file1 false
	// subdir/parent true
	// subdir/parent/child true
	// subdir/parent/child/file4 false
	// subdir/parent/file2 false
	// subdir/parent/file3 false

	return nil
}

/* SAÍDA ESPERADA PARA ESTE CÓDIGO
Listing subdir/parent
  child true
  file2 false
  file3 false
Listing subdir/parent/child
  file4 false
Visiting subdir
  subdir true
  subdir/file1 false
  subdir/parent true
  subdir/parent/child true
  subdir/parent/child/file4 false
  subdir/parent/file2 false
  subdir/parent/file3 false
*/
