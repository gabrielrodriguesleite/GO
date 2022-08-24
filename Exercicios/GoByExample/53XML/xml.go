package main

import (
	"encoding/xml"
	"fmt"
)

// Go vem com suporte imbutido para XML e formatos do tipo
// através do pacote encoding.xml

// Plant vai ser mapeada para XML. Do mesmo modo que no exemplo do JSON
// as tags dos os campos contém diretivas para codificação/decodificação.
// Aqui usamos algumas funções especiais do pacote XML: o campo XMLName
// dita do nome do elemento XML que representa a estrutura;
// id, attr, significa que o campo Id é um atributo XML ao invés de um
// elemento aninhado.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {

}
