package main

import "os"

// Um panic tipicamente é quando algo ocorre de uma forma inesperadamente errada
// Na maioria das vezes é usado para abortar a execução quando ocorrem erros que
// não deveriam de acontecer numa operação normal, ou que o código não está
// preparado para lidar.
func main() {

	// Panic será usado nos próximos exemplos para verificar por erros inesperados
	// Esta é a única aplicação nesses exemplos desenhada para disparar um panic

	// panic("a problem")

	// Um uso comum para panic é quando queremos abortar a aplicação se um função
	// retornar um erro que não sabemos ou não queremos lidar.
	// Aqui um exemplo de tratamento de erro quando uma função não conseguir
	// criar o arquivos gerando um erro que irá abortar a aplicação mostrando
	// o valor do erro.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// Rodando essa aplicação com a linha 14 descomentada irá fazer um panic,
// imprimindo a mensagem de erro na tela, traços da goroutine e sair
// com um status diferente de zero.

// A aplicação é abortada ao encontrar o primeiro panic por isso só irá executar
// a criação arquivo se a linha(14) com panic for comentada.

// Diferente de muitas linguagens onde se usam excessões para lidar com muitos
// tipos de erros, em Go isto é feito de uma maneira mais "idiomática" onde
// se retorna um valor de erro indicativo onde é possível.
