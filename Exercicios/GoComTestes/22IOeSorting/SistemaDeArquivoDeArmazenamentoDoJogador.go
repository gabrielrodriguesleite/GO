package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	// bancoDeDados io.ReadSeeker // https://golang.org/pkg/io/#ReadSeeker
	bancoDeDados io.Writer
	liga         Liga
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	return s.liga
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterPontuacaoDoJogador(nome string) int {
	jogador := s.liga.Buscar(nome)
	if jogador != nil {
		return jogador.Vitorias
	}
	return 0
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) GravarVitoria(nome string) {
	jogador := s.liga.Buscar(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		s.liga = append(s.liga, Jogador{nome, 1})
	}
	// s.bancoDeDados.Seek(0, 0)
	json.NewEncoder(s.bancoDeDados).Encode(s.liga)
}

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados io.ReadWriteSeeker) *SistemaDeArquivoDeArmazenamentoDoJogador {
	bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(bancoDeDados)
	return &SistemaDeArquivoDeArmazenamentoDoJogador{
		bancoDeDados: &fita{bancoDeDados},
		liga:         liga,
	}
}
