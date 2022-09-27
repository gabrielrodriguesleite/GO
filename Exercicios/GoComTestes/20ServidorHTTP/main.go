package main

import (
	"log"
	"net/http"
)

func main() {
	// tratador := http.HandlerFunc(ServidorJogador) // HandlerFunc é um adaptador
	// que permite usar funções comuns como tratadores.
	log.Println("Escutando na porta 5000")
	servidor := &ServidorJogador{NovoArmazenamentoJogadorEmMemoria()}
	// if err := http.ListenAndServe(":5000", tratador); err != nil {
	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}

func NovoArmazenamentoJogadorEmMemoria() *ArmazenamentoJogadorEmMemoria {
	return &ArmazenamentoJogadorEmMemoria{map[string]int{}}
}

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return a.armazenamento[nome]
}

func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	a.armazenamento[nome]++
}

// Utilizando curl para comprovar o funcionamento:
// curl -X POST http://localhost:5000/jogadores/Leite
// curl http://localhost:5000/jogadores/Leite
// deve retornar 1
