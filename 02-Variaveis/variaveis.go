package main

import "fmt"

func main()  {
	
	a := true
	b := 64
	c := 3.14
	d := 'A'
	e := "Palavra"
	f := `
  Um parágrafo
definido durante
a declaração e com
o tipo inferido pelo
valor.
`
	// vamos descobrir os tipos
	fmt.Println("\n VARIÁVEIS EM GO")
	fmt.Printf("O tipo de 'a' é %T e tem o valor de: %v\n",a,a)
	fmt.Printf("O tipo de 'b' é %T e tem o valor de: %v\n",b,b)
	fmt.Printf("O tipo de 'c' é %T e tem o valor de: %v\n",c,c)
	fmt.Printf("O tipo de 'd' é %T e tem o valor de: %v\n",d,d)
	fmt.Printf("O tipo de 'e' é %T e tem o valor de: %v\n",e,e)
	fmt.Printf("O tipo de 'f' é %T e tem o valor de: %v\n",f,f)
}