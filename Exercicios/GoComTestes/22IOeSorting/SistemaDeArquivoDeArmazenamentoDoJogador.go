package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	// bancoDeDados io.ReadSeeker // https://golang.org/pkg/io/#ReadSeeker
	bancoDeDados io.ReadWriteSeeker
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() Liga {
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
	jogador := liga.Buscar(nome)

	if jogador != nil {
		jogador.Vitorias++
	}
	s.bancoDeDados.Seek(0, 0)
	json.NewEncoder(s.bancoDeDados).Encode(liga)
}
