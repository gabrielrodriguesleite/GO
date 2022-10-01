package main

import (
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.ReadSeeker // https://golang.org/pkg/io/#ReadSeeker
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	s.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(s.bancoDeDados)
	return liga
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaPontuacaoDoJogador(nome string) int {
	return 0
}
