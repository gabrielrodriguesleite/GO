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

	// Forma e Parse ussando perfil baseado em exemplo
	// Normalmente usamos uma constante do pacote time para definir
	// o perfil, mas também é possível personalizar um perfil.
	p(t.Format("3:04PM"))
	// Perfis devem usar a referência "Mon Jan _2 15:04:05 2006"
	// para mostrar o padrão ao qual o format/parse deve formatar um time/string
	// O exemplo deve ter exatamente estes valores com mostrado.
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2021-08-10-03T14:00:04.000000-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)
}
