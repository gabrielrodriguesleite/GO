package main

import "fmt"

// go suporta estruturas e interfaces embebidas
// para expressar uma composição de tipos mais integrada(?)
// cuidado para não confundir com a diretiva //go:embed introduzida na versão 1.16+
// que embebe arquivos e pastas em um aplicativo binário

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// um container embebe uma base
// uma embebição parece como um campo sem um nome
type container struct {
	base
	str string
}

func main() {

	// quando criando estruturas com literais, temos que inicializa
	// a embebição explicitamente.
	// aqui o tipo embebido serve como o nome do amcpo
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// podemos acessar os campos de base diretamente. ex: co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// como também é possível chamar o nome completo
	fmt.Println("also num:", co.base.num)

	// como container embebe base, os métodos de base tornam-se métodos de container
	// da mesma forma podemos chamar diretamente o método embebido em co
	fmt.Println("describe:", co.describe())

	// ------------------

	type describer interface {
		describe() string
	}

	// estruturas embebidas com métodos podem ser usadas para conferir
	// implementação de interface em outras estruturas.
	// como podemos ver aqui container agora implementa a interface describer
	// por que ela embebe base
	var d describer = co
	fmt.Println("describer:", d.describe())
}
