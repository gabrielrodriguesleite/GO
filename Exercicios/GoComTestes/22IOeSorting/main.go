package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	// o_RDWR = abrir para leitura
	// o_CREATE = criar se não existir
	// 0666 permissões do arquivo vide: https://superuser.com/questions/295591/what-is-the-meaning-of-chmod-666
	if err != nil {
		log.Fatalf("problema abrindo %s %v", dbFileName, err)
	}

	armazenamento, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(db)
	if err != nil {
		log.Fatalf("problema criando o sistema de arquvio do armazenamento do jogador, %v ", err)
	}

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
