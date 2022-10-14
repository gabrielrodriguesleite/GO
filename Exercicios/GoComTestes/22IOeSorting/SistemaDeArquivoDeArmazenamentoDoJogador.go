package main

import (
	"encoding/json"
	"os"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	// bancoDeDados io.ReadSeeker // https://golang.org/pkg/io/#ReadSeeker
	bancoDeDados *json.Encoder
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

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(arquivo *os.File) *SistemaDeArquivoDeArmazenamentoDoJogador {
	arquivo.Seek(0, 0)
	liga, _ := NovaLiga(arquivo)
	return &SistemaDeArquivoDeArmazenamentoDoJogador{
		bancoDeDados: json.NewEncoder(&fita{arquivo}),
		liga:         liga,
	}
}
