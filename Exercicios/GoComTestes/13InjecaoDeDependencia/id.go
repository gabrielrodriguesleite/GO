package main

import (
	"fmt"
	"io"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Printf("Olá, %s", nome)
}
