package main

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
	GravarVitoria(nome string)
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
	roteador      *http.ServeMux
}

func NovoServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	s := &ServidorJogador{
		armazenamento,
		http.NewServeMux(),
	}
	s.roteador.Handle("/liga", http.HandlerFunc(s.manipulaLiga))
	s.roteador.Handle("/jogadores", http.HandlerFunc(s.manipulaJogadores))
	return s
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.roteador.ServeHTTP(w, r)
}

func (s *ServidorJogador) manipulaLiga(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *ServidorJogador) manipulaJogadores(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		s.processarVitoria(w, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(w, jogador)
	}
}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, jogador string) {
	pontuacao := s.armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) processarVitoria(w http.ResponseWriter, jogador string) {
	s.armazenamento.GravarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}
