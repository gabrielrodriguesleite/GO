package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes são frequentemente usados para computar identificadores curtos
// para binários ou partes de textos. Por exemplo, certificados TLS/SSL usam
// SHA256 para computar assinaturas de certificados. Neste exemplo temos uma
// forma de computar hashes SHA256 em Go.

// Go implementa várias funções hash nos pacotes crypto/*
func main() {

	s := "sha256 this string"

	// Aqui instanciamos um novo hash
	h := sha256.New()

	// Write espera bytes. Se temos uma string s, usamos
	// []byte(s) para transformar em bytes.
	h.Write([]byte(s))

	// Aqui recebemos a hash resultante como um slice de bytes.
	// O argumento de Sum pode ser usado para juntar a um slice de bytes existente.
	// Normalmente o argumento não é necessário.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
