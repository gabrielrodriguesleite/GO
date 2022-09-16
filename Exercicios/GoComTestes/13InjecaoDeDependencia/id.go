package main

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func main() {
	// Cumprimenta(os.Stdout, "Hello\n")
	err := http.ListenAndServe(
		":5000",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				Cumprimenta(w, "WEB")
			}))

	if err != nil {
		fmt.Println(err)
	}

}
