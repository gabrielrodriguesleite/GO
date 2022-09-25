package main

import (
	"fmt"
	"net/http"
)

func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):] // não é recomendável
	// mas serve por enquanto, para pegar o caminho da requisição.
	fmt.Fprint(w, ObterPontuacaoJogador(jogador))
}

func ObterPontuacaoJogador(nome string) string {
	if nome == "Leite" {
		return "20"
	}

	if nome == "Marcela" {
		return "25"
	}

	return ""
}

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}
