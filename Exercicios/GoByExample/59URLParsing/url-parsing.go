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

	// Acessando o esquema diretamente.
	fmt.Println(u.Scheme) // postgres

	// User contém todas as informações de autenticação.
	// Também é possível acessar os valores de nome e senha de forma individual.
	fmt.Println(u.User)            // user:pass
	fmt.Println(u.User.Username()) // user
	p, _ := u.User.Password()      //
	fmt.Println(p)                 // pass

	// Host contém  tanto o hostname e a porta se presentes.
	// Para extraí-los se usa SplitHostPort.
	fmt.Println(u.Host)                        // host.com:5432
	host, port, _ := net.SplitHostPort(u.Host) //
	fmt.Println(host)                          // host.com
	fmt.Println(port)                          // 5432

	// Extraindo o path e o fragmento que vem depois de #
	fmt.Println(u.Path)     // /path
	fmt.Println(u.Fragment) // f

	// Para acessar os parametros em uma string no formato k=v, usasse RawQuery.
	// Também é possível analisar parametros de consulta para um map.
	// O map de parametros de consulta analisado é de strings para slices de strings
	// então o index em [0] é para o caso de acessar apenas o primeiro valor.
	fmt.Println(u.RawQuery)            // k=v
	m, _ := url.ParseQuery(u.RawQuery) //
	fmt.Println(m)                     // map[k:[v]]
	fmt.Println(m["k"][0])             // v

}

/* SAÍDA ESPERADA PARA ESSE CÓDIGO
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
*/
