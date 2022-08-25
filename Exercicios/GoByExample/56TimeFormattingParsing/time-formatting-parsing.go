package main

import (
	"fmt"
	"time"
)

// Go suporta formatação e análise de formatos de tempo
// via layouts baseados em padrão.

func main() {

	p := fmt.Println

	// Este é um exemplo básico de formatação de tempo de acordo com
	// RFC3339, usando a constante do perfil correspondente.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Análise de tempo utiliza o mesmo valor de perfil como formato.
	t1, _ := time.Parse(time.RFC3339, "2021-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Wed Aug _2 21:39:05 2022"))
	p(t.Format("2021-08-10-03T14:00:04.000000-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)
}
