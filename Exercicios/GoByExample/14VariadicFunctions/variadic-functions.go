package main

import "fmt"

// função que recebe uma quantidade não definida de inteiros
// nums recebe o tipo []int e terá o len com o número de items
func sum(nums ...int) {

	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	// variadic functions podem receber um array ou slice espalhado
	sum(nums...)

}
