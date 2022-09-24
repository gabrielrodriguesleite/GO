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
		ctx := r.Context()
		data := make(chan string, 1)

		go func() { // fetch é executado e envia o resultado pelo canal
			data <- store.Fetch() // com 1 buffer ou seja se chegar antes
		}() // do ctx.Done é enviado se não apenas é cancelado.

		select { // toma ação com o que chegar primeiro, o dado do fetch
		case d := <-data: // ou o cancelamento do contexto.
			fmt.Fprint(w, d)
		case <-ctx.Done(): // sinal chamado quando o contexto finaliza.
			store.Cancel()
		}
	}
}

type Store interface {
	Fetch() string
	Cancel()
}
