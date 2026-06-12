# Companion — versão finalizada (backup)

Projeto do workshop **já completo**, após todos os passos do live code: o agente
`Gophi` (`character.go`) com Gemini 2.5 Flash + Google Search, servido por um chat
web em Go.

> Este é o estado final. No workshop, os participantes partem do esqueleto na raiz
> do repositório e criam o `character.go` do zero.

## Pré-requisitos

- Go 1.26+
- Uma `GOOGLE_API_KEY` do [Google AI Studio](https://aistudio.google.com/app/api-keys)

## Como rodar

```bash
export GOOGLE_API_KEY="sua-chave-do-ai-studio"
go run .
```

Acesse <http://localhost:5000> (no Cloud Shell, use o **Web Preview** na porta 5000).

## Variáveis de ambiente

| Variável | Obrigatória | Descrição |
|---|---|---|
| `GOOGLE_API_KEY` | sim | Chave do AI Studio usada pelo ADK para chamar o Gemini |
| `PORT` | não | Porta do servidor HTTP (default `5000`) |
| `PROJECT_ID` | não | Projeto GCP para Cloud Trace. Sem ele (ou sem credenciais), o tracing é desabilitado automaticamente e o app segue rodando |

## Estrutura

```
example/
├── main.go          # servidor HTTP + runner do ADK + tracing opcional
├── character.go     # agente Gophi (modelo, instrução, tools)
├── templates/
│   └── index.html   # UI do chat
└── static/
    ├── style.css
    ├── app.js       # typewriter, voz e lip sync
    └── images/      # imagens do personagem (boca aberta/fechada)
```
