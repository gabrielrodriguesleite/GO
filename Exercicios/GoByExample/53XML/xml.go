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

	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Primeiramente emitimos o XML representando a estrutura.
	// MarshalIdent produz uma saída de mais fácil compreensão
	out, _ := xml.MarshalIndent(coffee, " ", "  ") // → XML
	fmt.Println(string(out))                       // XML →

	// Para adicionar um header genérico para a saída
	fmt.Println(xml.Header + string(out))

	// Use Unmarshal para traduzir um fluxo de bytes com XML em uma estrutura de dados.
	// Se o XML estiver mal formado ou não puder ser mapeado na estrutura, um erro
	// descritivo será retornado.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// A etiqueta do campo parent>child>plant fala ao codificador para aninhar
	// todas as plants sob <parent><child>...
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

/* EXEMPLO DE SAÍDA DESTE CÓDIGO
<plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>

// -----------------

<?xml version="1.0" encoding="UTF-8"?>
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>

// -----------------

Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
 <nesting>
   <parent>
     <child>
       <plant id="27">
         <name>Coffee</name>
         <origin>Ethiopia</origin>
         <origin>Brazil</origin>
       </plant>
       <plant id="81">
         <name>Tomato</name>
         <origin>Mexico</origin>
         <origin>California</origin>
       </plant>
     </child>
   </parent>
 </nesting>
*/
