# Webhook Relay – Backend (Go)

## Visão Geral
Recebe webhooks e reenvia para um destino com retentativas exponenciais.

## Funcionalidades
- Endpoint `/hooks` para recepção
- Retentativa exponencial em falhas 5xx
- Preserva cabeçalhos

## Tecnologias Utilizadas
- net/http
- github.com/cenkalti/backoff

## Pré-requisitos
- Go 1.22+

## Como Executar
Clone o repositório:
```bash
git clone https://github.com/seu-usuario/webhook-relay.git
cd webhook-relay
```

Instale dependências e rode em dev:
```bash
go mod tidy
go run ./...
```

Acesse no navegador/cliente:
```text
POST http://localhost:8085/hooks
```

## Scripts Disponíveis
- `go run ./...` – executa em modo de desenvolvimento
- `go build` – gera binário de produção
- `go test ./...` – executa testes (se houver)

## Estrutura do Projeto
```
main.go
go.mod
```

## Habilidades Demonstradas
Integrações resilientes, políticas de retry, manipulação de headers.
