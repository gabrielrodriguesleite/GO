package main

import (
	"encoding/base64"
	"fmt"
)

// Go provê suporte integrado para codificação/decodificação base64
// Mais em: https://en.wikipedia.org/wiki/Base64

func main() {

	// A string que iremos codificar e decodificar
	data := "abc123!?$*&()'-=@~"

	// Go suporte tanto o padrão base64 quanto o tipo compatível com URL.
	// Aqui codificamos usando o tipo padrão.
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// A decodificação pode retornar um erro, pelo qual se pode verificar
	// se a entrada está no formato correto.
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	// URL

	// Exemplo usando base64 compatível com URL
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}

/* SAÍDA ESPERADA PARA O CÓDIGO
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~
YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
*/
