package main

import (
	"encoding/xml"
	"fmt"
)

// Go vem com suporte imbutido para XML e formatos do tipo
// através do pacote encoding.xml

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
