package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes são frequentemente usados para computar identificadores curtos
// para binários ou partes de textos. Por exemplo, certificados TLS/SSL usam
// SHA256 para computar assinaturas de certificados. Neste exemplo temos uma
// forma de computar hashes SHA256 em Go.

func main() {

	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
