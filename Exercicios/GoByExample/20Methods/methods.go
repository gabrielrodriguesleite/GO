package main

import "fmt"

type rect struct {
	width, height int
}

// go suporta métodos definidos em tipos estruturais

// metodos podem ser definidos por receber tipos ou ponteiros para estruturas
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {

	r := rect{width: 10, height: 5}

	// chamada dos métodos definidos para estrutura rect
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// go automaticamente manipula conversões entre valores e ponteiros para chamada de métodos
	// por receber um ponteiro se evita copiar na chamada dos métodos ou para permitir
	// que o método modifique a estrutura recebida
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
