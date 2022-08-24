package main

import "encoding/xml"

// Go vem com suporte imbutido para XML e formatos do tipo
// através do pacote encoding.xml

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func main() {

}
