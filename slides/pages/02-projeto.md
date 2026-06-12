---
layout: section
---

# O Projeto
## O que vamos construir

<!--
Antes de entrar nos conceitos e ferramentas, vale deixar claro o que vamos construir — assim a teoria que vem a seguir já faz sentido com um objetivo concreto em mente.
-->

---

# Companion — Visão Geral

<div class="abs-tr m-5 text-base font-bold" style="background:#eff6ff;border:1px solid #93c5fd;border-radius:9999px;padding:8px 18px;color:#1e40af">📦 github.com/o-mago/companion-workshop</div>

<div class="grid grid-cols-2 gap-8">
<div class="min-w-0">

<div class="mb-4">
<p class="font-bold mb-1">O que é</p>
<p class="text-sm text-slate-600">Um web app de chat com um agente de IA com personalidade — <strong>Gophi</strong>, um gopher Go hiperativo obcecado com goroutines.</p>
</div>

<div>
<p class="font-bold mb-2">Stack</p>
<ul class="text-sm text-slate-700 list-disc ml-4">
  <li>Go (net/http)</li>
  <li>Google ADK</li>
  <li>Gemini 2.5 Flash</li>
  <li>Templates HTML</li>
  <li>Cloud Shell / GCP</li>
</ul>
</div>

</div>
<div>

<div class="mb-4">
<p class="font-bold mb-2">Funcionalidades</p>
<ul class="text-sm text-slate-700 list-disc ml-4">
  <li>Chat em tempo real via HTTP</li>
  <li>Personalidade via system prompt</li>
  <li>Google Search integrado</li>
  <li>Servidor MCP externo (nano-banana)</li>
</ul>
</div>

**Interface**

```
┌--------------------------------------┐
│   Gophi 🐹                                   │
├---------------------------------------┤
│  Gophi: Oi! Já rodou go fmt?      │
│  Posso resolver com goroutines.│
│                                                     │
│  você: [____________________]   │
└--------------------------------------┘
```

</div>
</div>

<!--
O projeto é simples de propósito — o foco é nos conceitos, não em CSS ou infraestrutura. O main.go já está pronto com o servidor HTTP, rotas e runner. O que falta é o character.go com o agente.

Contextualize: ao final do live code vamos ter um chat funcional rodando no Cloud Shell, acessível via Web Preview, com um personagem que tem personalidade, acessa o Google em tempo real e pode gerar imagens.

Se alguém quiser levar para produção depois, o que muda é basicamente: trocar InMemoryService por um serviço de sessão persistente e colocar um load balancer na frente.
-->

<!-- --- -->

<!-- # Roteiro do Live Code

<div class="mt-2 space-y-3">

<div class="flex gap-4 items-start">
<div class="step-num">1</div>
<div class="text-sm pt-1"><strong>Criar o agente</strong> — <code>character.go</code> com <code>llmagent</code>, modelo Gemini 2.5 Flash e instrução básica</div>
</div>

<div class="flex gap-4 items-start">
<div class="step-num">2</div>
<div class="text-sm pt-1"><strong>Dar personalidade</strong> — refinar a instrução do Gophi: goroutines, GOPATH trauma e exemplos de resposta</div>
</div>

<div class="flex gap-4 items-start">
<div class="step-num">3</div>
<div class="text-sm pt-1"><strong>Adicionar Google Search</strong> — importar <code>geminitool</code> e habilitar busca em tempo real</div>
</div>

<div class="flex gap-4 items-start">
<div class="step-num">4</div>
<div class="text-sm pt-1"><strong>Conectar MCP</strong> — subir o <code>nano-banana-mcp</code> e configurar o Antigravity CLI</div>
</div>

<div class="flex gap-4 items-start">
<div class="step-num">5</div>
<div class="text-sm pt-1"><strong>Gerar assets com Antigravity CLI</strong> — prompt para gerar imagens e mover para <code>static/images</code></div>
</div>

</div>

<!--
Revele cada passo com clique. Isso cria ritmo e dá tempo para as pessoas visualizarem o que vem pela frente.

Steps 1–3 são os mais rápidos — cada um é uma modificação pequena no character.go. Step 4 envolve abrir um novo terminal e subir o MCP server. Step 5 é o mais "mágico" — um prompt gera as imagens e as move para o lugar certo.

Lembrete: cada passo tem um prompt pronto no README.md. Não precisa decorar nada — é só copiar, colar no Antigravity CLI e deixar ele trabalhar.

Tempo estimado por step: 5–8 minutos. Se ficar atrás, os steps 4 e 5 podem ser demonstrados pelo apresentador enquanto os outros acompanham.
-->

<!-- ---

# Estrutura do Projeto

```
companion-workshop/
├── main.go              # HTTP server, rotas, runner ADK
├── character.go         # Definição do agente (você vai criar!)
├── go.mod               # google.golang.org/adk
├── templates/
│   └── index.html       # UI do chat
├── static/
│   └── images/          # Assets gerados pelo Antigravity CLI
└── companion/           # Pacotes auxiliares
```

<div class="hl hl-slate mt-5">

**Variáveis de ambiente necessárias**

```bash
export GOOGLE_API_KEY="sua-chave-do-ai-studio"
# ou no Cloud Shell com Application Default Credentials:
gcloud auth application-default login
```

</div> -->

<!--
Mostre que a estrutura é intencionalmente simples. Não tem Makefile, Docker, CI — só o suficiente para o workshop funcionar.

O character.go não existe ainda — é o arquivo que vamos criar no Step 1 usando o Antigravity CLI. O main.go já tem um stub de `newRootAgent` que retorna nil; vamos substituir pelo agente real.

Sobre as variáveis de ambiente: o setup.sh que rodamos no Step 2 já exporta o GOOGLE_API_KEY automaticamente. Mas lembre as pessoas que cada novo terminal precisa ter a variável — por isso adicionamos ao ~/.bashrc.
-->
