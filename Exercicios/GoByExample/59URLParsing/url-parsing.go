package main

import (
	"fmt"
	"net"
	"net/url"
)

// URLs servem para prover uma forma uniforme para localizar recursos.
// Mais em: https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/
// Aqui alguns exemplos de como analizar URLs em Go.
func main() {

	// Neste exemplo será analizado uma URL de exemplo, que inclui um esquema,
	// informação de autenticação, hospedeiro, porta, caminho, parametros de consulta
	// e fragmento de consulta.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Analisando a URL e garantindo que não houveram erros.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Acessando o esquema diretamente
	fmt.Println(u.Scheme) // postgres

	// User contém todas as informações de autenticação
	// Também é possível acessar os valores de nome e senha de forma individual.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host contém  tanto o hostname e a porta se presentes.
	// Para extraí-los se usa SplitHostPort
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

}
