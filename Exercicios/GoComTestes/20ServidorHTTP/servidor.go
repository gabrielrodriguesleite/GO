package main

import (
	"fmt"
	"net/http"
)

func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):] // não é recomendável
	// mas serve por enquanto, para pegar o caminho da requisição.

	if jogador == "Leite" {
		fmt.Fprint(w, "20")
		return
	}

	if jogador == "Marcela" {
		fmt.Fprint(w, "25")
		return
	}

}
