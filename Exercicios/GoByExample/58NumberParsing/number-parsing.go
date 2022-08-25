package main

import (
	"fmt"
	"strconv"
)

// Analizar/interpretar números à partir de strings é uma tarefa
// tanto comum como básica para muitas aplicações.
// Aqui temos exemplos disso em Go:

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("567", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

}
