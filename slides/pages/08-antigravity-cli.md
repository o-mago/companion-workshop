---
layout: section
---

# Antigravity CLI

<!--
O Antigravity CLI vai ser nossa ferramenta de desenvolvimento durante o live code. Vale mostrar o que ele consegue fazer além do chat simples.
-->

---

# Antigravity CLI

<p class="text-slate-600 mb-5">Ferramenta de linha de comando open source do Google para interagir com modelos Gemini diretamente do terminal. Age como um agente de desenvolvimento.</p>

<div class="grid grid-cols-2 gap-8">
<div>

<p class="font-bold mb-2">Capacidades</p>
<ul class="text-sm text-slate-700 list-disc ml-4 space-y-1">
  <li>Chat interativo no terminal</li>
  <li>Lê e escreve arquivos do projeto</li>
  <li>Executa comandos shell</li>
  <li>Geração de código e imagens</li>
  <li>Suporte nativo a <strong>MCP servers</strong></li>
</ul>

</div>
<div>

<p class="font-bold mb-2">Como vamos usar</p>

<div class="space-y-2 text-sm">
<div class="hl hl-blue">
<strong>Pair programming</strong> — pedir ao CLI para criar e modificar arquivos Go durante o live code
</div>
<div class="hl hl-green">
<strong>MCP + nano-banana</strong> — gerar as imagens do personagem via tool call após configurar o servidor MCP
</div>
<div class="hl hl-slate">
<strong>Instalação</strong>
```bash
curl -fsSL https://antigravity.google/cli/install.sh | bash
agy  # inicia o agente
```
</div>
</div>

</div>
</div>

<!--
O Antigravity CLI é o diferencial deste workshop em relação a outros — ao invés de escrever código manualmente, usamos o próprio Gemini para escrever o código do agente Gemini. Meta e didático ao mesmo tempo.

O diferencial em relação a outros chats: acesso ao sistema de arquivos local. Ele pode ler main.go, entender o contexto do projeto, criar character.go no lugar certo, rodar `go build`, checar erros e corrigir — tudo em sequência. Não é só um chatbot.

Destaque o suporte nativo a MCP: depois de configurar o `mcp_config.json`, as tools ficam disponíveis automaticamente sem nenhum código adicional (lembre que o campo do servidor remoto é `serverUrl`).

Nota: o Gemini CLI antigo (npm `@google/gemini-cli`, comando `gemini`) foi descontinuado para contas individuais em 18/06/2026 — o caminho oficial agora é o Antigravity CLI (`agy`).
-->
