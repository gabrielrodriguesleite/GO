package main

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {
	t.Run("/liga de um leitor", func(t *testing.T) {

		bancoDeDados := strings.NewReader(`[
			{"Nome": "Leite", "Vitorias": 20},
			{"Nome": "Marcela", "Vitorias" : 25}
		]`)

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}
		recebido := armazenamento.PegaLiga()
		esperado := []Jogador{
			{"Leite", 20},
			{"Marcela", 25},
		}

		verificaLiga(t, recebido, esperado)
	})

	t.Run("pegar a pontuação do jogador", func(t *testing.T) {

		bancoDeDados := strings.NewReader(`[
			{"Nome": "Leite", "Vitorias": 20},
			{"Nome": "Marcela", "Vitorias" : 25}
		]`)

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}
		recebido := armazenamento.PegaPontuacaoDoJogador("Marcela")
		esperado := 25

		definePontuacaoIgual(t, recebido, esperado)
	})
}

func definePontuacaoIgual(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido %d esperado %d", recebido, esperado)
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
		os.remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}
