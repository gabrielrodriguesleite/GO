package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Escutando na porta 5000")
	servidor := NovoServidorJogador(NovoArmazenamentoJogadorEmMemoria())
	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}

// Utilizando curl para comprovar o funcionamento:
// curl -X POST http://localhost:5000/jogadores/Leite
// curl http://localhost:5000/jogadores/Leite
// deve retornar 1
