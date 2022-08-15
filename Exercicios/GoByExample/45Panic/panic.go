package main

// Um panic tipicamente é quando algo ocorre de uma forma inesperadamente errada
// Na maioria das vezes é usado para abortar a execução quando ocorrem erros que
// não deveriam de acontecer numa operação normal, ou que o código não está
// preparado para lidar.
func main() {

	// Panic será usado nos próximos exemplos para verificar por erros inesperados
	// Esta é a única aplicação nesses exemplos desenhada para disparar um panic
	panic("a problem")
}
