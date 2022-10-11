package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problema abrindo %s %v", dbFileName, err)
	}

	armazenamento := &SistemaDeArquivoDeArmazenamentoDoJogador{db}
	log.Println("Escutando na porta 5000")
	servidor := NovoServidorJogador(armazenamento)

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}

// Utilizando curl para comprovar o funcionamento:
// curl -X POST http://localhost:5000/jogadores/Leite
// curl http://localhost:5000/jogadores/Leite
// deve retornar 1

// https://larien.gitbook.io/aprenda-go-com-testes/criando-uma-aplicacao/io#refatore
