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
**Essa mudança só será aplicada depois do login ou por rodar o comando `source $HOME/.profile`**

---
###### referencia
https://go.dev/doc/install