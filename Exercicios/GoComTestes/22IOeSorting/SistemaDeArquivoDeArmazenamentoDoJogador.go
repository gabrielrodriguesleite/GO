package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	// bancoDeDados io.ReadSeeker // https://golang.org/pkg/io/#ReadSeeker
	bancoDeDados io.ReadWriteSeeker
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	s.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(s.bancoDeDados)
	return liga
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaPontuacaoDoJogador(nome string) int {
	jogador := s.PegaLiga().Buscar(nome)
	if jogador != nil {
		return jogador.Vitorias
	}
	return 0
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) SalvaVitoria(nome string) {
	liga := s.PegaLiga()

	for i, jogador := range liga {
		if jogador.Nome == nome {
			liga[i].Vitorias++ // pegar o valor por referência (e não pela cópia do elemento "jogador")
		}
	}
	s.bancoDeDados.Seek(0, 0)
	json.NewEncoder(s.bancoDeDados).Encode(liga)
}
