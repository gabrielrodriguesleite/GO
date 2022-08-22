package main

// Go possui suporte integrado para codificação e decodificação de JSON
// incluindo de e para tipos de dados integrados e personalizados.

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

}
