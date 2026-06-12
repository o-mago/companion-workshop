---
layout: section
---

# Ferramentas do Workshop

<!--
Antes de entrar na teoria, vale alinhar o ambiente. Estas ferramentas são o que vamos usar durante todo o workshop — garantir que todos saibam o que é cada uma evita confusão na hora do live code.
-->

---

# GCP — Google Cloud Platform

<div class="grid grid-cols-2 gap-8 mt-2">
<div>

<p class="text-sm text-slate-600 mb-5">Plataforma de computação em nuvem do Google. Oferece infraestrutura, serviços gerenciados e ferramentas de IA/ML usadas por empresas de todos os tamanhos.</p>

<p class="font-bold mb-2">O que usamos no workshop</p>
<ul class="text-sm text-slate-700 list-disc ml-4 space-y-1">
  <li><strong>Cloud Shell</strong> — ambiente de desenvolvimento no browser</li>
  <li><strong>APIs</strong> — habilitação de serviços Google</li>
</ul>

</div>
<div>

<p class="font-bold mb-2">Por que GCP no workshop?</p>

<div class="space-y-3">
<div class="hl hl-green text-sm">
<strong>Zero instalação local</strong> — tudo roda no browser, sem configurar Go, Git ou credenciais na máquina
</div>
<div class="hl hl-blue text-sm">
<strong>Ambiente padronizado</strong> — todos os participantes com o mesmo setup, sem "funciona na minha máquina"
</div>
<div class="hl hl-slate text-sm">
<strong>Integração nativa</strong> — autenticação, APIs e ferramentas Google já conectadas
</div>
</div>

</div>
</div>

<!--
Muitos participantes podem já conhecer AWS ou Azure — contextualize GCP como o equivalente Google. Não precisa aprofundar.

O ponto mais importante aqui é o motivo de usar GCP: zero atrito de setup. Em workshops presenciais, configurar ambiente local consome 30-40% do tempo. Com Cloud Shell isso some.
-->

---

# Cloud Shell

<div class="grid grid-cols-2 gap-8 mt-2">
<div>

<p class="text-sm text-slate-600 mb-4">Ambiente de desenvolvimento completo rodando diretamente no browser, hospedado pelo Google. Sem instalação, sem configuração local.</p>

<p class="font-bold mb-2">O que já vem incluso</p>
<ul class="text-sm text-slate-700 list-disc ml-4 space-y-1">
  <li>Go, Python, Node.js, Java</li>
  <li><code>gcloud</code> CLI autenticado</li>
  <li>Git, Docker, Vim, VS Code (Web)</li>
  <li>5 GB de armazenamento persistente</li>
  <li><strong>Web Preview</strong> — expõe portas locais via HTTPS</li>
</ul>

</div>
<div>

<p class="font-bold mb-2">Web Preview</p>
<p class="text-sm text-slate-600 mb-3">Permite acessar serviços rodando no Cloud Shell pelo browser — é assim que vamos acessar nosso chat.</p>

```bash
# Aplicação rodando na porta 5000
go run .

# URL gerada automaticamente:
# https://5000-cs-XXXX.cs-region.cloudshell.dev
```

<div class="hl hl-orange text-sm mt-3">
Acesse em <strong>console.cloud.google.com</strong> → ícone do terminal no topo direito
</div>

</div>
</div>

<!--
Demonstre o Cloud Shell ao vivo se possível — abrir o terminal e mostrar que Go já está instalado causa um bom impacto.

Web Preview é o recurso mais importante para o workshop: é o que permite testar a aplicação sem deploy. Mostre onde fica o botão de Web Preview na interface do Cloud Shell.

Lembrete: o Cloud Shell hiberna após inatividade e o processo Go morre. Se alguém perder a conexão, basta rodar `go run .` de novo.
-->

---

# Google AI Studio

<div class="grid grid-cols-2 gap-8 mt-2">
<div>

<p class="text-sm text-slate-600 mb-4">Plataforma web para explorar, testar e integrar modelos Gemini via API. É o ponto de entrada para desenvolvedores que querem usar IA generativa sem infraestrutura.</p>

<p class="font-bold mb-2">Principais recursos</p>
<ul class="text-sm text-slate-700 list-disc ml-4 space-y-1">
  <li>Playground interativo para prompts</li>
  <li>Geração de API Keys</li>
  <li>Visualização de uso e limites</li>
  <li>Exportar código pronto (Go, Python, JS)</li>
  <li>Fine-tuning de modelos</li>
</ul>

</div>
<div>

<div class="space-y-3">

<div class="hl hl-blue text-sm">
<strong>API Key</strong> — chave simples para autenticar chamadas. É o que geramos no Step 1 do workshop e exportamos como <code>GOOGLE_API_KEY</code>.
</div>

<div class="hl hl-green text-sm">
<strong>Free tier generoso</strong> — Gemini 2.5 Flash disponível gratuitamente com limites de RPM e TPD suficientes para desenvolvimento e workshops.
</div>

<div class="hl hl-slate text-sm">
<strong>aistudio.google.com</strong> — acesso direto, sem precisar de projeto GCP ou billing habilitado para começar.
</div>

</div>

</div>
</div>

<!--
AI Studio é onde os participantes criaram a API Key no Step 1. Vale abrir ao vivo e mostrar o playground — permite testar prompts antes de escrever código, o que é uma boa prática de desenvolvimento.

Destaque a diferença fundamental com o Gemini Enterprise Agent Platform: AI Studio usa API Key simples, o Gemini Enterprise Agent Platform usa credenciais de projeto GCP com billing. Para o workshop usamos AI Studio justamente para eliminar essa fricção.
-->

---

# Gemini Enterprise Agent Platform

<div class="grid grid-cols-2 gap-8 mt-2">
<div>

<p class="text-sm text-slate-600 mb-4">Plataforma enterprise do Google para ML e IA generativa. Integrada à GCP com billing, IAM, VPC e SLAs de produção.</p>

<p class="font-bold mb-2">Recursos principais</p>
<ul class="text-sm text-slate-700 list-disc ml-4 space-y-1">
  <li>Todos os modelos Gemini via API</li>
  <li>Agentes gerenciados com UI</li>
  <li>Grounding com Google Search e dados próprios</li>
  <li>Integração com BigQuery, Cloud Storage</li>
  <li>Avaliação e monitoramento de agentes</li>
</ul>

</div>
<div>

<p class="font-bold mb-2">AI Studio vs Gemini Enterprise Agent Platform</p>

<div class="text-sm mt-2">

| | AI Studio | Gemini Enterprise |
|---|---|---|
| Auth | API Key | ADC / Service Account |
| Billing | Gratuito (limites) | Pay-per-use |
| SLA | Não | Sim |
| VPC / IAM | Não | Sim |
| Uso ideal | Dev / workshops | Produção |

</div>

</div>
</div>

<!--
O Gemini Enterprise Agent Platform é o destino natural depois do workshop — quando o projeto sair do laptop e virar produção, a migração é basicamente trocar a autenticação.

O ADK em Go suporta os dois backends: basta remover GOOGLE_GENAI_USE_VERTEXAI ou setar para TRUE. A API de código não muda.

Agent Platform (antes Vertex AI Agents) é a oferta managed — você sobe um agente sem escrever servidor HTTP. Vale mencionar para quem não quer gerenciar infraestrutura.
-->

