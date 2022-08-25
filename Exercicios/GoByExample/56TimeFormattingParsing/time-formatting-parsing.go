package main

import (
	"fmt"
	"time"
)

// Go suporta formatação e tradução de formatos de tempo
// via layouts baseados em padrão

func main() {

	p := fmt.Println

	// Este é um exemplo básico de formatação de tempo de acordo com
	// RFC3339, usando a constante layout correspondente
	t := time.Now()
	p(t.Format(time.RFC3339))
}
