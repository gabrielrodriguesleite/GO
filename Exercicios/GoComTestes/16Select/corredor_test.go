package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("testa se retorna a url mais rápida", func(t *testing.T) {
		servidorLento := criaServidorComAtraso(20 * time.Millisecond)
		servidorRapido := criaServidorComAtraso(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida
		resultado, _ := Corredor(URLLenta, URLRapida)

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidorA := criaServidorComAtraso(11 * time.Second)
		servidorB := criaServidorComAtraso(12 * time.Second)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Corredor(servidorA.URL, servidorB.URL)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})
}

func criaServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}
