package main

import (
	"context"
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
		// Será utilzado o context para que não seja necessário
		// chamar uma função de cancelamento pra cada método.
		// O handler será responsável em emitir o contexto à
		// Store em cascata e tratar o erro que virá da Store
		// quando essa for cancelada.
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Dessa forma o servidor ficou simplificado pois não é mais responsavel
// pelo cancelamento. Ele simplesmente passa o context e confia nas
// funções em cascata, respeitando qualquer cancelamento que possa ocorrer.
