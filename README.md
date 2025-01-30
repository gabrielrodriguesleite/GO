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
1. Para uma instalação por usuário incluir no arquivo* $HOME/.profile:
2. Para uma instalação no sistema incluir no arquivo /etc/profile

`export PATH=$PATH:/usr/local/go/bin`

**Essa mudança só será aplicada depois do login ou por rodar o comando**

`source $HOME/.profile`

> *(Nota) Se estiver utilizando zsh o arquivo $HOME/.zshrc também precisa receber o comando `export PATH=$PATH:/usr/local/go/bin`

Para verificar a instalação:

`go version`

A versão instalada deve aparecer ex: 

`go version go1.18.4 linux/amd64`

O comando `go env` mostra as variaveis de ambiente definidas.

###### Referências

https://go.dev/doc/install

---

## Hello World em Go

`helloworld.go`

```go
package main // define o escopo da aplicação ajudando na separação de responsabilidade
import "fmt" // a biblioteca padrão que trás diversas facilidades

func main() { // define uma nova função 
	fmt.Println("Hello World") // utiliza um método da biblioteca padrão para imprimir a mensagem na tela

  // Mais sobre a biblioteca padrão em:
  // https://pkg.go.dev/fmt
}
```

O comando `go run helloworld.go` irá executar o arquivo.

Deve exibir `Hello World` na saída padrão.

### 🎁 BÔNUS - web server rodando no docker

`helloweb.go`

```go
package main

import (
  "fmt"
  "net/http"
  "time"
)

func greet(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
  http.HandleFunc("/", greet)
  http.ListenAndServe(":8080", nil)
}

```

#### 🐋 Docker

```sh
docker run --volume $(pwd):/go -it --rm -p 8080:8080 --name golang-dev docker.io/golang:1.23

go run helloweb.go
```

#### 🛜 Curl

Em outro terminal.

```sh
curl http://localhost:8080
```

---

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

---
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

Constantes são definidas no escopo da função usando apenas "=" ao invés de ":=" ou utilizando `const`

```go
const b int = 8
const s strint = "let's go"
const (
  pi = 3.14
  c int = 25
)
```

---
## Go mod e Go work - Módulos e Espaço de trabalho
Por padrão nossos projetos precisam ficar dentro da pasta `$HOME/go/src` para que possamos acessar um pacote dentro de outro arquivo.

Porém para podermos criar e manter projetos complexos precisamos ter uma maior flexibilidade com respeito a árvore de arquivos do nosso projeto.

Para isso existe o Go mod.

Dentro da pasta que desejamos trabalhar nosso módulo executamos `go mod init nome_do_modulo`

Um arquivo `go.mod` será criado, esse arquivo define nosso modulo e indica a versão do go utilizada no momento da criação, também registra as dependências do módulo.

Depois de instalarmos alguma dependência com o `go get nome_do_pacote` além do registro no `go.mod` ser atualizado também é criado um arquivo `go.sum` que serve pra travar a versão da dependência.

**Dicas**
O comando `go mod tidy` faz uma verificação e remove dependências não utilizadas também como instala dependências que estão faltando.

O comando `go mod graph` mostra todas as dependências do projeto. 

O comando `go mod vendor` cria uma pasta e inclui as dependências junto ao projeto.

O comando `go install` é o comando correto para instalar binários.

Considere o seguinte exemplo:

./main.go

./exemplo/exemplo.go

```sh
# Definir um go.mod partir da pasta do módulo
cd exemplo
go mod init exemplo

# Definir um go.work (workspace) ou seja um arquivo com as dependencias locais
cd ../
go work init
go work use exemplo
go run main.go
```

./main.go

./go.work

./exemplo/exemplo.go

./exemplo/go.mod

O arquivo `go.work` criado na raiz do projeto indica o caminho dos módulos usados pelo nosso pacote `main`.

###### Referências
https://go.dev/doc/tutorial/workspaces
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md

---

## Escopo e visibilidade

O escopo em Go é bem definido.

Declarações globais, feitas fora de uma função, podem ser acessadas dentro do pacote.

Funções declaradas dentro de um pacote estarão visiveis a outros pacotes se forem declaradas começando com a letra maiúscula. 

Nomes de funções começados com letras minusculas são visiveis apenas dentro do mesmo pacote. 

O mesmo se aplica para os atributos de uma _struct_.

./04-Escopo/pacote/pacote.go
```go
package pacote

func Imprimir() {

}
```

./04-Escopo/main.go
```go
package main

import "pacote"

func main() {
  pacote.Imprimir()
}
```

A função Imprimir só estará disponível pois inicia com 'I' maiúsculo.

## Go by example

Exercícios sobre diversos pacotes embutidos resolvendo tarefas comuns estão na pasta:

```
Exercicios/GoByExample
```

###### Referências

https://gobyexample.com/


## Aprendendo Go com testes

Exercício de teste onde a linguagem é explicada por meio de testes estão na pasta:

```sh
Exercicios/GoComTestes
```

###### Referências

https://larien.gitbook.io/aprenda-go-com-testes/

---

## Rodando testes em go

```sh
go test ./...                   # Executa todos os testes de forma recursiva (todos os subdiretórios)
go test -count=1 ./...          # Desabilita o cache do teste e indica que rode 1 vez
go test -v -count=10 ./...      # Roda 10 vezes de forma verbosa (mostra informações úteis)
go test -v -count=1 -race ./... # Teste condições de "data race", opção também disponível ao executar a aplicação `go run race .`
```

###### Referências
[Dicas rodando testes em Golang](https://www.youtube.com/watch?v=7pu7zXXYmCo)

---

## Dicas úteis:
Para quem está acostumado a desenvolver API com node utilizando nodemon para auxiliar nos carregamentos automatizados aí vai a dica:
1. Instale nodemon no sistema (dependência: npm): `npm i -g nodemon`
2. Rode a aplicação go utilizando nodemon com `nodemon --exec "go run main.go" --signal SIGTERM`
3. Bônus: para facilitar é possível usar tanto _Makefile_ quanto _package.json script_ para o passo 2.
