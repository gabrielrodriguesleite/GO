package main

import (
	"fmt"
	"path/filepath"
	"strings"
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

	// Dir e Base podem ser usados para separar um caminho em pasta e arquivo.
	// De maneira alternativa, Split irá retornar ambos na mesma chamada:
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	f, d := filepath.Split(p)
	fmt.Printf("Split(p) → Dir: %s, Base: %s\n", f, d)

	// É possível também testar se um caminho é relativo ou absoluto.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Alguns arquivos possuem nomes com a extenção indicada após o ponto.
	// É possível dividir a extenção e capturar ela com Ext:
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))
}
