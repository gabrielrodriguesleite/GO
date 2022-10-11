package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	// bancoDeDados io.ReadSeeker // https://golang.org/pkg/io/#ReadSeeker
	bancoDeDados io.ReadWriteSeeker
	liga         Liga
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	s.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(s.bancoDeDados)
	return liga
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterPontuacaoDoJogador(nome string) int {
	jogador := s.ObterLiga().Buscar(nome)
	if jogador != nil {
		return jogador.Vitorias
	}
	return 0
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) GravarVitoria(nome string) {
	liga := s.ObterLiga()
	jogador := liga.Buscar(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		liga = append(liga, Jogador{nome, 1})
	}
	s.bancoDeDados.Seek(0, 0)
	json.NewEncoder(s.bancoDeDados).Encode(liga)
}

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados io.ReadWriteSeeker) *SistemaDeArquivoDeArmazenamentoDoJogador {
	bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(bancoDeDados)
	return &SistemaDeArquivoDeArmazenamentoDoJogador{
		bancoDeDados: bancoDeDados,
		liga:         liga,
	}
}
