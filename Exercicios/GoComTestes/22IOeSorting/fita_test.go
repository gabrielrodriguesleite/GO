package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFita_Escrita(t *testing.T) {
	arquivo, limpa := criaArquivoTemporario(t, "12345")
	defer limpa()

	fita := &fita{arquivo}
	fita.Write([]byte("abc"))

	arquivo.Seek(0, 0)
	novoConteudoDoArquivo, _ := ioutil.ReadAll(arquivo)

	recebido := string(novoConteudoDoArquivo)
	esperado := "abc"

	if recebido != esperado {
		t.Errorf("recebido '%s' esperado '%s'", recebido, esperado)
	}

}

func criaArquivoTemporario(t *testing.T, dadoInicial string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	arquivotmp, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("não foi possível escrever o arquivo temporário %v", err)
	}

	arquivotmp.Write([]byte(dadoInicial))

	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}
