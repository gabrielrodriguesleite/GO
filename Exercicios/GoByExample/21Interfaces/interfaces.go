package main

import (
	"fmt"
	"math"
)

// interfaces são coleções nomeadas de assinaturas de methodos
// exemplo básico de uma interface para formas geometricas
type geometry interface {
	area() float64
	perim() float64
}

// implementamos interfaces para os tipos rect e circle
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// para implementar uma interface em go, precisamos apenas
// implementar todos os métodos na interface.
// aqui implementamos geometry em rects
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

// aqui implementamos geometry em circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// se uma variável tem o tipo de uma interface, então podemos
// chamar os métodos que existem na interface nomeada
// aqui a função generica measure se aproveita disso para funcionar
// com qualquer estrutura que implementam a interface
func measure(g geometry) {

	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {

	// tanto circle como rect implementam a interface geometry então
	// é possível usá-los como argumentos para measure
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

}
