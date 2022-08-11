package main

import "fmt"

// a partir da versão 18 go incluiu suporte para generics
// também conhecido como parametros de tipo

// um exemplo de função genérica, MapKeys recebe um map
// de quaquer tipo e retorna um slice das suas chaves
// Esta função tem 2 parametros typo - K e V
// K tem o comparable constraint, segnifica que podemos usar
// os operadores == e != para comparar seus valores.
// Isto é necessário para chaves de map em go.
// V tem a constraint any, significa que não é restrito de nenhuma forma
// any é um apelido para interface{}
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// como um exemplo do tipo generics,
// List é um tipo de lista unida com valores de qualquer tipo
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// é possível definir métodos no tipo generic como em tipos
// normais, mas é preciso manter o parametro tipo no lugar
// neste caso o tipo é List[T] e não List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys m:", MapKeys(m))

	// quando invocamos funções de generics podemos nos basear em inferencia de tipo
	// neste caso não precisamos especificar os pipos de K e V quando chamamos MapKeys
	// já que o compilador será capaz de inferir automaticamente
	// mas é possível sim especificar explicitamente

	// a marcação de warning é por que os tipos podem ser omitidos pois serão
	// definidos por inferência
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

}
