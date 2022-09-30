package main

import (
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
		recebido := armazenamento.("Marcela")
		esperado := 25

		if recebido != esperado {
			t.Errorf("recebido %d esperado %d", recebido, esperado)
		}
	})
}
