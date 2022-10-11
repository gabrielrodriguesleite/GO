package main

import (
	"fmt"
	"net/http"
)

// func ServidorJogador(w http.ResponseWriter, r *http.Request) {
// 	jogador := r.URL.Path[len("/jogadores/"):] // não é recomendável
// 	// mas serve por enquanto, para pegar o caminho da requisição.
// 	fmt.Fprint(w, ObterPontuacaoDoJogador(jogador))
// }

// func ObterPontuacaoDoJogador(nome string) string {
// 	if nome == "Leite" {
// 		return "20"
// 	}

// 	if nome == "Marcela" {
// 		return "25"
// 	}

// 	return ""
// }

type ArmazenamentoJogador interface {
	ObterPontuacaoDoJogador(nome string) int
	RegistrarVitoria(nome string)
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		s.registrarVitoria(w, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(w, jogador)

	}
}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, jogador string) {
	pontuacao := s.armazenamento.ObterPontuacaoDoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) registrarVitoria(w http.ResponseWriter, jogador string) {
	s.armazenamento.RegistrarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}
