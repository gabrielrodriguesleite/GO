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
