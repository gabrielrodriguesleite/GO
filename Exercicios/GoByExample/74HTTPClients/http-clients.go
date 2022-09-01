package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// A biblioteca padrão de go vem com um suporte excelente para clientes e
// servidores HTTP no pacote "net/http".
// Neste exemplo este pacote será usado para lançar uma simple requisição HTTP
func main() {

	// Emitir uma requisição HTTP GET para um server usando http.Get é um atalho
	// conveniente para criar um http.Client e chamar o método Get.
	// Isto utiliza o objeto http.DefaultClient que possui configurações padrões úteis.
	res, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Imprimimos o código de estado da resposta.
	fmt.Println("Response status:", res.Status)

	// Em seguida imprimimos as 5 primeira linhas do corpo da resposta
	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
