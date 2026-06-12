---
layout: section
---

# Go no Ecossistema de IA

<!--
Uma dúvida comum: "por que Go para IA se todo mundo usa Python?" Vamos endereçar isso diretamente.
-->

---

# O que Go Suporta Hoje

<div class="grid grid-cols-2 gap-8 mt-1">
<div>

**SDKs Oficiais**

<div class="text-sm">

| Pacote | Descrição |
|--------|-----------|
| `google.golang.org/adk` | Agent Development Kit |
| `google.golang.org/genai` | Gemini / Gemini Enterprise |
| `github.com/modelcontextprotocol/go-sdk` | MCP server/client |

</div>

</div>
<div>

<div class="card card-orange text-sm">

**Ponto de atenção**

O ADK em Go já está **estável**. A versão Python é mais madura, mas o Go já tem:

- `LLMAgent`, `Runner`, `Session`
- Built-in tools (GoogleSearch)
- Custom function tools
- Streaming de eventos via `range`
- Multi-agent

</div>

</div>
</div>

<div class="hl hl-blue mt-2 text-sm">

**Por que Go?** Performance, binários pequenos, tipagem forte, excelente para servidores HTTP — ideal para agentes em produção na nuvem.

</div>

<!--
Seja honesto sobre o estado do ADK em Go: é estável. Isso significa que a API é confiável, a documentação está disponível e os recursos essenciais para construir um agente funcional já estão lá.

A escolha de Go faz sentido quando o agente precisa ser um serviço de produção: startup rápido, memória baixa, binário único sem dependências, tipagem forte que pega bugs em compile time.

Python domina no mundo de ML/treino de modelos, mas para servir um agente como API HTTP, Go tem vantagens reais. O ecossistema está crescendo rápido — SDKs oficiais do Google já existem.
-->
