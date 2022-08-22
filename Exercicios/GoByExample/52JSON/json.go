package main

import (
	"encoding/json"
	"fmt"
)

// Go possui suporte integrado para codificação e decodificação de JSON
// incluindo de e para tipos de dados integrados e personalizados.

// Estas duas estruturas serão utilizadas para demonstrar codificação e
// decodificação de tipos personalizados
type response1 struct {
	Page   int
	Fruits []string
}

// Apenas campos exportados vão ser codificados/decodificados em JSON.
// Campos devem iniciar com letras maiúsculas para serem exportados.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Primeiramente vamos ver codificação de tipos básicos para strings JSON.
	// Aqui alguns exemplos para valores atômicos
	bolB, _ := json.Marshal(true) // → json
	fmt.Println(string(bolB))     // json →

	intB, _ := json.Marshal(1) // → json
	fmt.Println(string(intB))  // json →

	fltB, _ := json.Marshal(3.14) // → json
	fmt.Println(string(fltB))     // json →

	strB, _ := json.Marshal("gopher") // → json
	fmt.Println(string(strB))         // json →

	slcD := []string{"maçã", "pêssego", "pêra"}
	slcB, _ := json.Marshal(slcD) // → json
	fmt.Println(string(slcB))     // json →
}
