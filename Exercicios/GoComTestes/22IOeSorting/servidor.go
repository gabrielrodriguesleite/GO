package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const tipoDoConteudoJSON = "application/json"

type ArmazenamentoJogador interface {
	ObterPontuacaoDoJogador(nome string) int
	GravarVitoria(nome string)
	ObterLiga() Liga
}

type Jogador struct {
	Nome     string
	Vitorias int
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
	http.Handler  // incorporação do Handler dentro do ServidorJogador para acesso
	// direto das funções.
	// Mais em: https://golang.org/doc/effective_go.html#embedding
}

func NovoServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	s := new(ServidorJogador)
	s.armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/liga", http.HandlerFunc(s.ManipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(s.ManipulaJogadores))

	s.Handler = roteador
	return s
}

func (s *ServidorJogador) ManipulaLiga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", tipoDoConteudoJSON)     // Essa ordem
	json.NewEncoder(w).Encode(s.armazenamento.ObterLiga()) // importa
	w.WriteHeader(http.StatusOK)
}

func (s *ServidorJogador) ManipulaJogadores(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		s.processarVitoria(w, jogador)
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

func (s *ServidorJogador) processarVitoria(w http.ResponseWriter, jogador string) {
	s.armazenamento.GravarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}
