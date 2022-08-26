package main

import (
	"encoding/base64"
	"fmt"
)

// Go provê suporte integrado para codificação/decodificação base64
// Mais em: https://en.wikipedia.org/wiki/Base64

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

}
