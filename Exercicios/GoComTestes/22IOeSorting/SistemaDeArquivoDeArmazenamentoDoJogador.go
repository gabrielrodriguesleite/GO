package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.Reader
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	var liga []Jogador
	json.NewDecoder(s.bancoDeDados).Decode(&liga)
	return liga
}
