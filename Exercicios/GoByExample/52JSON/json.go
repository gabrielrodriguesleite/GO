package main

import (
	"encoding/json"
	"fmt"
)

// Go possui suporte integrado para codificação e decodificação de JSON
// incluindo de e para tipos de dados integrados e personalizados.

// Estas duas estruturas serão utilizadas para demonstrar codificação e
// decodificação de tipos personalizados
type response1 struct {
	Page   int
	Fruits []string
}

// Apenas campos exportados vão ser codificados/decodificados em JSON.
// Campos devem iniciar com letras maiúsculas para serem exportados.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// Regra: Chaves iniciando em maiúscula será codificada
// o nome da chave codificada pode ser modificado utilizando o recurso de tag

func main() {

	// Primeiramente vamos ver codificação de tipos básicos para strings JSON.
	// Aqui alguns exemplos para valores atômicos
	bolB, _ := json.Marshal(true) // → json
	fmt.Println(string(bolB))     // json →

	intB, _ := json.Marshal(1) // → json
	fmt.Println(string(intB))  // json →

	fltB, _ := json.Marshal(3.14) // → json
	fmt.Println(string(fltB))     // json →

	strB, _ := json.Marshal("gopher") // → json
	fmt.Println(string(strB))         // json →

	// Aqui alguns exemplos de codificação/decodificação de slices e maps.
	slcD := []string{"maçã", "pêssego", "pêra"}
	slcB, _ := json.Marshal(slcD) // → json
	fmt.Println(string(slcB))     // json →

	mapD := map[string]int{"maçã": 5, "alface": 7}
	mapB, _ := json.Marshal(mapD) // → json
	fmt.Println(string(mapB))     // json →

	// -------------

	// O pacote JSON pode codificar automaticamente seus tipos de dados personalizados.
	// Ele irá incluir apenas campos exportados na saída codificada e vai por padrão
	// usar seus nomes como chaves JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"maçã", "pêssego", "pêra"},
	}
	res1B, _ := json.Marshal(res1D) // → json
	fmt.Println(string(res1B))      // json →

	// É possível usar etiquetas nos campos da declaração da estrutura para personalizar
	// os nomes das chaves codificadas em JSON. Verifique a definição de response2 para
	// ver os exemplos de etiquetas deste tipo.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"maçã", "pêssego", "pêra"},
	}
	res2B, _ := json.Marshal(res2D) // → json
	fmt.Println(string(res2B))      // json →

	// ----------

	// Decodificando dados JSON em valores GO.
	// Aqui um exemlo de uma estrutura de dados genérica
	byt := []byte(`{"num":3.14, "strs":["a", "b"]}`)

	// É preciso providenciar uma variável onde o pacote JSON possa colocar os dados
	// decodificados. Este map vai conter um map de strings para tipos de dados arbitrários.
	var dat map[string]interface{}

	// Aqui a decodificação e o tratamento de erro
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) // json →

	// Para podermos usar o dado decodificado é necessário converter para um tipo apropriado.
	// Neste exemplo o valor da chave num é convertido para um tipo experado float64
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

}
