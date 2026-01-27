# BFF Relaunch - Backend For Frontend

Um **BFF (Backend For Frontend)** desenvolvido em Go que atua como gateway de API, recebendo requisições do frontend e orquestrando a comunicação com os microsserviços do ecossistema Relaunch.

## 🎯 Sobre o Projeto

Este projeto é responsável por centralizar e gerenciar todas as requisições vindas do frontend, distribuindo as responsabilidades entre diferentes microsserviços através de comunicação gRPC. Ele implementa autenticação JWT, validação de requisições e fornece uma API REST padronizada para o cliente.

## 🚀 Tecnologias

- **Go** - Linguagem de programação principal
- **Gin** - Framework web para criação da API REST
- **gRPC** - Comunicação entre microsserviços
- **JWT** - Autenticação e autorização
- **Protocol Buffers** - Serialização de dados

## 🏗️ Arquitetura

O BFF atua como uma camada intermediária que:

- Recebe requisições HTTP/REST do frontend
- Valida tokens JWT e autentica usuários
- Roteia as requisições para os microsserviços apropriados via gRPC
- Agrega respostas de múltiplos serviços quando necessário
- Retorna dados formatados para o cliente

## 📋 Funcionalidades Principais

- **Autenticação e Autorização**: Middleware JWT para validação de tokens
- **Gerenciamento de Usuários**: CRUD de usuários
- **Notificações**: Sistema de notificações em tempo real
- **Validação de Requisições**: Validação robusta de dados de entrada
- **Tratamento de Erros**: Respostas padronizadas e informativas

## 🔐 Segurança

- Validação de tokens JWT em todas as rotas protegidas
- Suporte a múltiplos formatos de claims (camelCase e snake_case)
- Sanitização de headers de autorização
- Configuração de secrets via variáveis de ambiente
