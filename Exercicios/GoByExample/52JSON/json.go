package main

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

}
