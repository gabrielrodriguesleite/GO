package main

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
// depois criar uma implementação que dá suporte ao mecanismo preferido.
