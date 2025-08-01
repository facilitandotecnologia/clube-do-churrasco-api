# Documento de Requisitos do Produto (PRD)  
**Produto:** Clube do Churrasco  
**Objetivo:** Organizar churrascos com amigos

---

## 1. Visão Geral

O Clube do Churrasco é um aplicativo web/mobile projetado para ajudar usuários a organizar eventos de churrasco com amigos. O app facilita a criação de eventos, gestão de convidados, solicitações e confirmações de pagamento, gerenciamento de contas de usuário e consulta de preços em casas de carne.

---

## 2. Funcionalidades

### 2.1. Gerenciamento de Conta de Usuário (CRUD)
- Usuários podem **Criar** uma conta com informações obrigatórias (nome, e-mail, número do WhatsApp, senha).
- Usuários podem **Visualizar** seus dados e perfil.
- Usuários podem **Atualizar** suas informações (ex: alterar nome, e-mail, número do WhatsApp, senha).
- Usuários podem **Excluir** sua conta, removendo todos os dados pessoais e associações a eventos.

### 2.2. Criar Evento
- Usuários podem criar um novo evento de churrasco.
- Campos obrigatórios:  
  - **Data:** Quando o evento acontecerá.  
  - **Descrição:** Resumo do evento.  
  - **Participantes:** Lista de amigos convidados.  
  - **Detalhes do Evento:** Observações ou descrição adicional.

### 2.3. Adicionar Amigos ao Evento
- Usuários podem buscar e adicionar amigos ao evento como participantes.
- Amigos recebem notificações/convites.

### 2.4. Consultar Preços em Casas de Carne
- Usuários podem pesquisar ou consultar preços de carnes em casas de carne locais.
- Integração com APIs externas ou entrada manual de preços pelas casas de carne.

### 2.5. Solicitar Pagamento aos Participantes
- O administrador do evento pode solicitar pagamento de cada participante para cobrir os custos do evento.
- Solicitações de pagamento são enviadas individualmente para cada participante.
- Participantes recebem solicitações de pagamento pelo app e/ou via mensagem no WhatsApp.
- O status do pagamento é acompanhado por participante.

### 2.6. Participante Envia Comprovante de Pagamento & Reconhecimento em Tempo Real
- Participantes podem enviar o comprovante de pagamento (ex: recibo ou boleto) para o administrador via app.
- O app processa e reconhece o pagamento em tempo real (ex: via integração com gateways de pagamento ou análise automatizada do comprovante).
- O administrador é notificado instantaneamente quando um pagamento é reconhecido e confirmado.

---

## 3. Histórias de Usuário

- **Como usuário,** quero criar, visualizar, atualizar e excluir minha conta para gerenciar minhas informações pessoais e privacidade.
- **Como usuário,** quero criar um evento de churrasco para organizar encontros com meus amigos.
- **Como usuário,** quero adicionar amigos ao meu evento para que possam ser convidados e participar.
- **Como usuário,** quero consultar preços de carnes em lojas locais para planejar minhas compras e orçamento.
- **Como administrador do evento,** quero solicitar pagamento dos participantes para dividir os custos do evento.
- **Como participante,** quero receber solicitações de pagamento pelo app ou WhatsApp para saber quanto e como pagar.
- **Como participante,** quero enviar meu comprovante de pagamento pelo app para que meu pagamento seja reconhecido e confirmado em tempo real.
- **Como administrador,** quero ser notificado instantaneamente quando o pagamento de um participante for reconhecido.

---

## 4. Modelo de Dados

- **Usuário**
  - id
  - nome
  - e-mail
  - número do WhatsApp
  - senha (hash)
- **Evento**
  - id
  - data
  - descrição
  - detalhes
  - participantes (lista de IDs de usuários)
  - administrador (ID do usuário)
  - solicitaçõesDePagamento (lista de objetos de solicitação de pagamento)
- **CasaDeCarne**
  - id
  - nome
  - localização
  - precosDeCarne (lista de tipos de carne e preços)
- **SolicitaçãoDePagamento**
  - id
  - eventoId
  - participanteId
  - valor
  - status (pendente, pago, vencido)
  - enviadoPor (app, WhatsApp)
  - comprovanteDePagamento (arquivo/imagem ou ID da transação)
  - reconhecido (booleano)
  - reconhecidoEm (timestamp)

---

## 5. Requisitos Não Funcionais

- Design responsivo para web e mobile.
- Autenticação segura para usuários.
- Atualizações em tempo real para convites, solicitações de pagamento, reconhecimento de pagamento e alterações de preços.
- Integração com API do WhatsApp para envio de solicitações de pagamento.
- Integração com gateways de pagamento ou análise de comprovantes para reconhecimento em tempo real.
- Exclusão de dados conforme LGPD para contas de usuário.

---

## 6. Fora do Escopo

- Processamento de pagamento (transferência real de fundos).
- Gestão de entrega ou logística.

---