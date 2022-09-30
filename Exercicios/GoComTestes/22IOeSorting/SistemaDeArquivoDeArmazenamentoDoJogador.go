package main

import (
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.Reader
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	liga, _ := NovaLiga(s.bancoDeDados)
	return liga
}
