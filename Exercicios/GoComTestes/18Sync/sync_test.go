package main

import (
	"sync"
	"testing"
)

// PROBLEMÁTICA: O objetivo é criar um contador seguro para concorrência.
// Primeiramente será desenvolvido um contador não seguro para verificar
// o funcionamento em um ambiente com apenas 1 linha de execução.
// Então serão utilizados várias go routines para expor a falha e então
// consertar.
func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, contador, 3)
	})

	t.Run("roda concorrentemente em seguranca", func(t *testing.T) {
		contagemEsperada := 1000
		contador := Contador{}

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incrementa()
				w.Done()
			}(&wg)
		}
		wg.Wait()
		verificaContador(t, contador, contagemEsperada)
	})
}

func verificaContador(t *testing.T, resultado Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
	}
}

//  + sync.WaitGroup é uma maneira simples de sincronizar processos concorrentes
/* Um "WaitGroup" aguarda por uma coleção de goroutines terminar seu processamento.
A goroutine principal faz a chamada para o "Add" definir o número de goroutines que
serão esperadas. Então, cada uma das goroutines é executada e chama "Done" quando
termina a execução. Ao mesmo tempo, "Wait" pode ser usado para bloquear a execução
até que todas as goroutines tenham terminado.
*/

// Com a implementação atual o teste falha as vezes já que pode acontecer de mais
// de uma goroutine tentar alterar o valor do contador ao mesmo tempo.
