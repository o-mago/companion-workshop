---
layout: section
---

# ADK
## Agent Development Kit

<!--
Chegamos no framework que vai sustentar o projeto. ADK é o que cola tudo junto — modelo, ferramentas, sessão e loop do agente.
-->

---

# O que é o ADK?

<p class="text-slate-600 mb-6"><strong>Google Agent Development Kit</strong> — framework open source para construir agentes de IA de produção.</p>

<div class="grid grid-cols-3 gap-5">
<div class="card card-green">

**Multi-linguagem**

Python, **Go**, TypeScript, Java e Kotlin

`google.golang.org/adk`

</div>
<div class="card card-blue">

**Abstrações prontas**

LLMAgent, Runner, Session, Tools — você foca na lógica

</div>
<div class="card card-purple">

**Integrações**

Gemini, Gemini Enterprise Agent Platform, Google Search, MCP servers

</div>
</div>

<!--
ADK é open source — o repositório está no GitHub. Vale mencionar isso porque a comunidade está crescendo e PRs são bem-vindos.

Python é a versão mais madura com mais exemplos e docs. O ADK em Go já é estável (GA) e tem tudo que precisamos para o workshop. Além de Python e Go, o ADK também suporta TypeScript, Java e Kotlin. Se alguém vier da comunidade Go e quiser contribuir, é uma ótima oportunidade.

O ADK 2.0 (mai/2026) introduziu o Workflow Runtime — execução baseada em grafo no lugar do executor hierárquico. Os docs agora vivem em adk.dev.

O valor principal do ADK: você não precisa implementar o loop do agente, gerenciamento de sessão, ou o protocolo de tool use. Tudo isso já está lá.
-->
