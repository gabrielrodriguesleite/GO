package main

import (
	"encoding/json"
	"fmt"
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
	s.bancoDeDados.Encode(s.liga)
}

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(arquivo *os.File) (*SistemaDeArquivoDeArmazenamentoDoJogador, error) {
	arquivo.Seek(0, 0)
	liga, err := NovaLiga(arquivo)
	if err != nil {
		return nil, fmt.Errorf("problema carregando o armazenamento do jogador de arquivo %s, %v", arquivo.Name(), err)
	}

	return &SistemaDeArquivoDeArmazenamentoDoJogador{
		bancoDeDados: json.NewEncoder(&fita{arquivo}),
		liga:         liga,
	}, nil
}
