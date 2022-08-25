package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O pacote math/rand provê geração de números pseudo-aleatórios
// Mais em: https://en.wikipedia.org/wiki/Pseudorandom_number_generator

func main() {

	// Este exemplo imprime 2 números aleatórios inteiros entre 0 e 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// Este exemplo retorna um ponto flutuante de 0.0 à f < 1.0
	fmt.Println(rand.Float64())

	// Também é possível gerar um ponto flutuante em outra variação
	// por exemplo 5.0 <= f < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// ---------------------- SEED

	// Por padrão o gerador é deterministico, então ele vai produzir a mesma
	// sequência de numeros cada vez que for executado.
	// Para produzir sequências variadas, forneça uma semente que muda.
	// Esta implementação não é uma forma segura de gerar números aleatórios
	// para senhas. Para essa tarefa use crypto/rand.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Chamar o rand.Rand resultante da mesma forma como chamaria as outras
	// funções do pagote rand.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// ---------------------- SEED FOOT SHOT

	// Caso seja usada a mesma semente, o gerador irá produzir a mesma
	// sequência de números aleatórios XD.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)

	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()

}
