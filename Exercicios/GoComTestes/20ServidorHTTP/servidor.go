package main

import (
	"fmt"
	"net/http"
)

// func ServidorJogador(w http.ResponseWriter, r *http.Request) {
// 	jogador := r.URL.Path[len("/jogadores/"):] // não é recomendável
// 	// mas serve por enquanto, para pegar o caminho da requisição.
// 	fmt.Fprint(w, ObterPontuacaoJogador(jogador))
// }

// func ObterPontuacaoJogador(nome string) string {
// 	if nome == "Leite" {
// 		return "20"
// 	}

// 	if nome == "Marcela" {
// 		return "25"
// 	}

// 	return ""
// }

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	pontuacao := s.armazenamento.ObterPontuacaoJogador(jogador)
	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	// w.WriteHeader(http.StatusNotFound) // implementar o mínimo para ot teste passar vai
	// evidenciar as lacunas no teste.
	// Mostrará que não estamos validado se é retornado um status ok caso o jogador seja
	// encontrado.

	fmt.Fprint(w, pontuacao)
}
