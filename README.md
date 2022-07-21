# GO
Repositorio sobre o estudo de Go Lang

## Instalação
Baixar a última versão: https://go.dev/dl/go1.18.4.linux-amd64.tar.gz

Remover instalações prévias da pasta `/usr/local/go`

`sudo rm -rf /usr/local/go`

Descompactar o arquivo baixado em `/usr/local`

`sudo tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz`

**Importante**
Não descompactar o arquivo dentro de uma instalação existente!

Incluir /usr/local/go/bin no ambiente PATH (apenas 1 opção)
1. Para uma instalação por usuário incluir no arquivo $HOME/.profile:
2. Para uma instalação no sistema incluir no arquivo /etc/profile

`export PATH=$PATH:/usr/local/go/bin`

**Essa mudança só será aplicada depois do login ou por rodar o comando**

`source $HOME/.profile`

Para verificar a instalação:

`go version`

A versão instalada deve aparecer ex: 

`go version go1.18.4 linux/amd64`

O comando `go env` mostra as variaveis de ambiente definidas.

---
###### Referências
https://go.dev/doc/install

## Hello World em Go

~/helloworld.go
```go
package main // define o escopo da aplicação ajudando na separação de responsabilidade
import "fmt" // a biblioteca padrão que trás diversas facilidades

func main() { // define uma nova função 
	fmt.Println("Hello World") // utiliza um método da biblioteca padrão para imprimir a mensagem na tela
}
```

O comando `go run helloworld.go` irá executar o arquivo.

Deve exibir `Hello World` na saída padrão.

## Compilando com build

Para compilarmos o pacote utilizamos `go build helloworld.go`

Com isso é criado um novo arquivo executável para a arquitetura atual. 

Para executar o arquivo recém criado no linux usamos o comando `./helloworld` 

A mensagem `Hello World` deve ser exibida na saída padrão.

Para compilarmos para outra arquitetura podemos definir a variável de ambiente `GOOS`

`GOOS=linux go build helloworld.go`

`GOOS=windows ARCH=386 go build helloworld.go`

`GOOS=darwin go build helloworld.go`

Para exibir uma lista completa use `go tool dist list`

Para exibir uma ajuda sobre o comando build use `go help build`

## Variaveis

```go
// vaiáveis declaradas e atribuídas em qualquer contexto
var x int
x = 256
var y, z string = "Olá", "Mundo"

func main() {
  w := 16 // para atribuir durante a declaração é necessário estar dentro do escopo da função
}
```

  ./02-Variaveis/variaveis.go
```go	
//...
a := true       // bool
b := 64         // int
c := 3.14       // float64
d := 'A'        // int32
e := "Palavra"  // string
f := `
  Um parágrafo
definido durante
a declaração e com
o tipo inferido pelo
valor.
`               // string

// outros tipos estão disponíveis mas esses são os mais comuns

// vamos descobrir os tipos
fmt.Println("\n VARIÁVEIS EM GO")
fmt.Printf("O tipo de 'a' é %T e tem o valor de: %v\n",a,a)
fmt.Printf("O tipo de 'b' é %T e tem o valor de: %v\n",b,b)
fmt.Printf("O tipo de 'c' é %T e tem o valor de: %v\n",c,c)
fmt.Printf("O tipo de 'd' é %T e tem o valor de: %v\n",d,d)
fmt.Printf("O tipo de 'e' é %T e tem o valor de: %v\n",e,e)
fmt.Printf("O tipo de 'f' é %T e tem o valor de: %v\n",f,f)
//...
```