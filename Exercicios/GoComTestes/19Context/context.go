package main

import (
	"fmt"
	"net/http"
)

// CONTEXTO

/* O pacote context ajuda a gerenciar processos demorados
de forma consistente dentro da aplicação.

Por exemplo quando a aplicação está buscando um recurso e
no meio recebe um sinal para desistir. Simplesmente abandonar
a goroutine e deixá-la processando não é nem eficiente nem ágil.

*/

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
}
