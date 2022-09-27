package main

import (
	"log"
	"net/http"
)

func main() {
	// tratador := http.HandlerFunc(ServidorJogador) // HandlerFunc é um adaptador
	// que permite usar funções comuns como tratadores.
	log.Println("Escutando na porta 5000")
	servidor := &ServidorJogador{&ArmazenamentoJogadorEmMemoria{}}
	// if err := http.ListenAndServe(":5000", tratador); err != nil {
	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}

type ArmazenamentoJogadorEmMemoria struct{}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return 123
}

func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {}
