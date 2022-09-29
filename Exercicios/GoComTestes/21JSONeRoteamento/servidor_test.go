package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Desafio parte 2
// criar um novo endpoint chamado /liga que retorne uma lista contendo
// todos os jogadores armazenados como um JSON.

// ===== parte 1
// Desafio: Criar um servidor web para que os usuários
// possam acompanhar quantas partidas os jogadores venceram.

// GET /jogadores/{nome} deve retornar o numero total de vitórias

// POST /jogadores/{nome} deve registrar uma vitória pra este nome de
// jogador, incrementando a cada nova chamada de submissão de dados.

// ==================================================================

// Abordagem: Desenvolvimento Orientado a Testes:
// criar um software que funciona o mais rápido possível, e a cada ciclo
// realizar pequenas melhorias até ter uma solução completa.

// Essa abordagem tras algumas vantagems:
// - Escopo de problema pequeno em todos os momentos.
// - Poucos detalhes ajudam a manter o foco.
// - Possibilidade de voltar o código em caso de bloqueio do desenvolvimento.

// ==================================================================

// --- PASSOS DO TDD ---
// 1. Escrever um teste
// 2. Ver o teste falhar
// 3. Escrever a menor quantidade de código para fazer o teste passar, mesmo
//    que usando código errado. (Isto é importante pois expõe testes falhos.
// 		[falsos positivos])
// 4. Reescrita do teste. (Refatoração)

// ==================================================================

// Mas como desenvolver pequenas partes desse projeto que necessita do
// post pra incluir dados para que possam ser recuperados com get?
// A resposta é INTERFACES que vão ser úteis quando precisarmos
// mockar (enganar) nossa função para fazer os testes no ambiente simulado.

// - GET precisa de um armazenamento para obter as informações. Isso deve
// ser uma interface para que ao executar os teste seja possível criar um
// código simples sem precisar criar um código final complexo de armazenamento.

// - POST, é preciso inspecionar as chamadas feitas ao armazenamento para ter
// certeza que os dados foram enviados corretamente. A implementação de
// gravação não ficará vinculada à busca dos dados.

// Para ter um código funcionando rapidamente, o ideal é fazer uma
// implementação simples da interface necessária(armazenamento) para só
// depois criar uma implementação que dá suporte ao mecanismo preferido.]

func TestLiga(t *testing.T) {

	t.Run("retorna a tabela da Liga como JSON", func(t *testing.T) {
		ligaEsperada := []Jogador{
			{"Leite", 32},
			{"Marcela", 35},
			{"Pedro", 10},
		}
		armazenamento := EsbocoArmazenamentoJogador{nil, nil, ligaEsperada}
		servidor := NovoServidorJogador(&armazenamento)

		requisicao := novaRequisicaoDeLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		obtido := obterLigaDaResposta(t, resposta.Body)
		verificaRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificaLiga(t, obtido, ligaEsperada)
		if resposta.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("resposta não tinha o tipo de conteúdo de application/json, obtido %v", resposta.Result().Header)
		}
	})

	t.Run("retorna 200 em /liga", func(t *testing.T) {
		armazenamento := EsbocoArmazenamentoJogador{}
		servidor := NovoServidorJogador(&armazenamento)
		requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
		responsta := httptest.NewRecorder()

		servidor.ServeHTTP(responsta, requisicao)

		var obtido []Jogador

		err := json.NewDecoder(responsta.Body).Decode(&obtido)

		if err != nil {
			t.Fatalf("Não foi possível fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", responsta.Body, err)
		}

		verificaRespostaCodigoStatus(t, responsta.Code, http.StatusOK)
	})
}

// ==================== TESTES PARTE 1 ====================

// ==================== TESTE requisições ====================
func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Leite":   20,
			"Marcela": 25,
		},
		nil, nil,
	}
	servidor := NovoServidorJogador(&armazenamento)
	t.Run("retornar resultado de Leite", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Leite")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificaCorpoRequisicao(t, resposta.Body.String(), "20")
	})

	t.Run("retonrar resultado para Marcela", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Marcela")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificaCorpoRequisicao(t, resposta.Body.String(), "25")
	})

	t.Run("retorna 404 para jogador não encontrado", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Jorge")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Code
		esperdado := http.StatusNotFound

		if recebido != esperdado {
			t.Errorf("recebido status %d esperado %d", recebido, esperdado)
		}
	})
}

// ==================== TESTE armazenamento ====================
func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil, nil,
	}
	servidor := NovoServidorJogador(&armazenamento)

	t.Run("registra vitorias na chamada ao método HTTP POST", func(t *testing.T) {
		jogador := "Marcela"

		requisicao := novaRequisicaoRegistrarVitoriaPost(jogador)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaRespostaCodigoStatus(t, resposta.Code, http.StatusAccepted)

		if len(armazenamento.registrosVitorias) != 1 {
			t.Errorf("verifiquei %d chamadas a RegistrarVitoria, esperava %d", len(armazenamento.registrosVitorias), 1)
		}

		if armazenamento.registrosVitorias[0] != jogador {
			t.Errorf("não registrou o vencedor corretamente, recebi '%s', esperava '%s'", armazenamento.registrosVitorias[0], jogador)
		}
	})
}

// ==================== TESTE integração ====================
func TestRegistrarVitoriasEBuscarEstasVitorias(t *testing.T) {
	armazenamento := NovoArmazenamentoJogadorEmMemoria()
	servidor := NovoServidorJogador(armazenamento)
	jogador := "Leite"

	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))

	resposta := httptest.NewRecorder()
	servidor.ServeHTTP(resposta, novaRequisicaoObterPontuacao(jogador))
	verificaRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

	verificaCorpoRequisicao(t, resposta.Body.String(), "3")
}

// ==================== FUNÇÕES AUXILIARES ====================

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificaCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisição é inválido, obtive '%s', esperava '%s'", recebido, esperado)
	}
}

func verificaRespostaCodigoStatus(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("não recebeu código de status HTTP esperado, recebido %d, esperado %d", recebido, esperado)
	}
}

type EsbocoArmazenamentoJogador struct {
	pontuacoes        map[string]int
	registrosVitorias []string
	liga              []Jogador
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func (e *EsbocoArmazenamentoJogador) GravarVitoria(nome string) {
	e.registrosVitorias = append(e.registrosVitorias, nome)
}

func (e *EsbocoArmazenamentoJogador) ObterLiga() []Jogador {
	return e.liga
}

func novaRequisicaoRegistrarVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func obterLigaDaResposta(t *testing.T, body io.Reader) (liga []Jogador) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&liga)
	if err != nil {
		t.Fatalf("Não foi possível fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", body, err)
	}
	return
}

func verificaLiga(t *testing.T, obtido, esperado []Jogador) {
	t.Helper()
	if !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("obtido %v esperado % v", obtido, esperado)
	}
}

func novaRequisicaoDeLiga() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/liga", nil)
	return req
}
