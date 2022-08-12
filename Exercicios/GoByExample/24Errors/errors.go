package main

import (
	"errors"
	"fmt"
)

// go comunica erros diretamente por um retorno explícito separado.
// go permite o fácil reconhecimento das funções que retornam erro
// como também uma manipulação fácil usando os recursos comuns da linguagem

func f1(arg int) (int, error) {
	if arg == 42 {
		// por convenção erros devem ser o último item retornado e
		// tem o tipo error como interface
		// errors.New constroi um erro básico com a mensagem passada
		return -1, errors.New("can't work with 42")
	}
	// quando não há erros retornamos um 'nil'
	return arg + 3, nil
}

// é possivel customizar o erro implementando o metodo Error na
// estrutura criada
type argError struct {
	arg  int
	prob string
}

// nesse exemplo é usado um tipo customizado pra representar um erro
// de argumento explicitamente
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// nesse caso usamos a sintaxe &argError pra construir uma nova
// estrutura passando os campos necessários
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil

}

func main() {

	// os 2 loops testam cada função de retorno de erro
	for _, i := range []int{7, 42} {
		// teste se houve erro escrito em linha com é de costume em go
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// se for necessario usar os dados em um erro customizado é possível
	// pegar o erro como uma instancia do tipo erro customizado por meio de asserção
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
