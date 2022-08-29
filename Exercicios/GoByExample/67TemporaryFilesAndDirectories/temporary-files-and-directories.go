package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Durante toda execução do programa, muitas vezes queremos
// criar dados que não são necessários após a finalização do
// programa. Arquivos e diretórios temporários são úteis para
// este propósito desde que não poluam o sistema de arquivos
// com o passar do tempo.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// A maneira mais fácil de crirar um arquivo temporário em Go
	// é chamando os.CreateTemp. Este método cria e abre um arquivo
	// para leitura e escrita. Passando "" no primeiro argumento
	// fará com que o arquivo seja criado no local padrão para seu OS.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Mostrando o nome e a localização do arquivo temporário.
	// Em sistemas Unix o diretório base deve ser "/tmp".
	// O nome do arquivo leva o prefixo definido no segundo argumento,
	// o restante do nome é definido automaticamente de forma a garantir
	// que ao chamar esse método novamente crie um nome diferente.
	fmt.Println("Temp file name:", f.Name())

	// Removemos o arquivo depois do fim da excução do programa. O próprio
	// sistema se encarrega de limpar os arquivos temporários de tempos em
	// tempos, mas é uma boa prática fazer isso explicitamente.
	defer os.Remove(f.Name())

	// -----------

	// Escrevendo alguns dados no arquivo.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// No caso de termos muitos arquivos temporários, e preferível utilizarmos
	// um diretório temporário. Os argumentos de os.MkdirTemp funcionam da mesma
	// forma que CreateTemp, porém retorna um nome de diretório ao invés de um
	// arquivo aberto.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	// RemoveAll apaga toda árvore de diretórios e arquivos dentro do diretório
	// passado como argumento.
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}
