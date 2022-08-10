package main

import "fmt"

func main() {

	// um slice
	nums := []int{2, 3, 4}
	sum := 0

	// range usado para somar numeros de um slice. (também funcionaria com array)
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)
}
