package main

import (
	"fmt"
	"path/filepath"
)

// O pacote filepath possui funções para analisar e contruir caminhos de arquivos
// de forma portável entre sistemas operacionais. Por exemplo dir/file no linux
// e dir\file no windows.
func main() {

	// Join deve ser usado quando se constroem paths de forma portável.
	// Aceita n número de argumentos e constroe a hierarquia do caminho à partir destes.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// É importante sempre dar preferência ao Join ao invés de concatenar caminhos
	// manualmente com /s ou \s. Além de prover portabilidade, Join também normaliza
	// os caminhos por remover separadores e mudanças de pastas desnecessárias.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

}
