# Notas de Explicação — Companion Workshop

Referência de código para uso durante o live code. Cada seção explica um arquivo do projeto.

---

## `main.go` — Servidor HTTP e Runner do ADK

```go
package main
```

Ponto de entrada da aplicação. Contém o servidor HTTP, as rotas e a inicialização do runner do ADK.

---

### Variáveis globais

```go
var (
    agentRunner *runner.Runner
    tmpl        *template.Template
)
```

- `agentRunner` — instância do runner do ADK. Começa como `nil`; só é inicializado se um agente válido for criado. Isso permite rodar a aplicação sem `character.go` (o agente fica em modo echo).
- `tmpl` — template HTML da interface de chat, carregado uma vez na inicialização.

---

### Constantes

```go
const (
    appName   = "companion"
    userID    = "inapp_user"
    sessionID = "default_session"
)
```

- `appName` — identifica a aplicação no ADK (usado internamente na sessão).
- `userID` e `sessionID` — fixos para simplificar o workshop. Em produção, cada usuário e conversa teria IDs únicos para isolar o histórico de mensagens.

---

### `handleIndex`

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

Rota `GET /`. Renderiza o `templates/index.html` e devolve o HTML da interface de chat. Não recebe dados do usuário — só serve a página.

---

### `handleChat`

```go
func handleChat(w http.ResponseWriter, r *http.Request) {
```

Rota `POST /chat`. É aqui que o frontend manda a mensagem e recebe a resposta do agente.

**Decodificação do body:**
```go
var req chatRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "bad request", http.StatusBadRequest)
    return
}
```
Lê o JSON `{ "message": "..." }` do corpo da requisição.

**Modo echo (sem agente):**
```go
if agentRunner == nil {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(chatResponse{Response: req.Message})
    return
}
```
Se `character.go` ainda não foi criado, o servidor funciona em modo espelho — devolve a mensagem do usuário sem processar. Permite testar a interface antes do Step 1.

**Execução do agente:**
```go
msg := genai.NewContentFromText(req.Message, genai.RoleUser)

var sb strings.Builder
for event, err := range agentRunner.Run(r.Context(), userID, sessionID, msg, agent.RunConfig{}) {
```
O ADK devolve um iterador de eventos. `range` é a sintaxe Go para iterar sobre esse stream — o loop continua até o agente terminar de processar (incluindo chamadas a ferramentas).

```go
    if event.IsFinalResponse() && event.Content != nil {
        for _, p := range event.Content.Parts {
            sb.WriteString(p.Text)
        }
    }
```
Filtra só o evento de resposta final. Eventos intermediários (chamadas a tools, raciocínio) existem mas são ignorados aqui. O texto de todos os parts é concatenado no `strings.Builder`.

---

### `newRootAgent` (stub)

```go
func newRootAgent(_ context.Context) (agent.Agent, error) {
    return nil, nil
}
```

Placeholder que existe para o código compilar antes do `character.go` ser criado. No Step 1 do live code, o Gemini CLI vai substituir esta função pela implementação real importada do `character.go`.

---

### `main`

```go
rootAgent, err := newRootAgent(ctx)
```
Tenta criar o agente. Se `newRootAgent` devolver `nil, nil` (stub), a aplicação sobe sem runner.

```go
sessionSvc := session.InMemoryService()

agentRunner, err = runner.New(runner.Config{
    AppName:           appName,
    Agent:             rootAgent,
    SessionService:    sessionSvc,
    AutoCreateSession: true,
})
```
`InMemoryService` armazena o histórico de conversa na memória do processo — simples e sem dependências. `AutoCreateSession: true` cria a sessão automaticamente se não existir ainda.

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /", handleIndex)
mux.HandleFunc("POST /chat", handleChat)
mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
```
Três rotas: página inicial, endpoint de chat e arquivos estáticos (imagens, CSS, JS). `http.StripPrefix` remove o prefixo `/static/` antes de buscar o arquivo no disco.

---

## `character.go` — Definição do Agente

Este arquivo **não existe no repositório** — é criado durante o live code no Step 1 usando o Gemini CLI.

```go
package main
```

Mesmo pacote que `main.go`, o que permite que a variável `rootAgent` seja acessada diretamente.

---

### Criação do modelo

```go
m, err := gemini.NewModel(ctx, "gemini-2.5-flash", nil)
```

Instancia o modelo Gemini 2.5 Flash via ADK. O terceiro parâmetro (nil) aceita configurações extras como temperatura e top-p — deixamos como padrão.

---

### Configuração do agente

```go
return llmagent.New(llmagent.Config{
    Model:       m,
    Name:        "companion_agent",
    Instruction: `...`,
    Tools:       []tool.Tool{ geminitool.GoogleSearch{} },
})
```

- `Model` — o LLM que processa as mensagens.
- `Name` — identificador interno do agente, usado em logs e no histórico de sessão.
- `Instruction` — o system prompt. É a peça mais importante: define quem o agente é, como fala e o que pode ou não fazer.
- `Tools` — lista de ferramentas disponíveis. `GoogleSearch` permite ao agente buscar informações em tempo real antes de responder.

---

### O system prompt do waku

```
You are waku, a witty, super sweet, and super intelligent cat.
```

Cada linha tem uma função específica:

| Linha | Função |
|---|---|
| Definição de persona | Estabelece quem é o agente — nome, espécie, personalidade |
| Core Rules | Restrições de comportamento — o que nunca fazer |
| Favorite Topics | Domínio de conhecimento e estilo de interação |
| Example Response Style | Few-shot examples — mostra o padrão de resposta esperado |
| `Answer no more than 3 sentences` | Controla o tamanho da resposta |

Os exemplos de resposta são fundamentais: quando é difícil descrever em palavras o tom exato que se quer, mostrar um exemplo ensina o modelo por imitação.

---

## `static/app.js` — Frontend do Chat

JavaScript responsável por toda a interatividade da interface: enviar mensagens, animação lip sync e síntese de voz.

---

### Referências ao DOM

```js
const characterImage = document.getElementById('character-image');
const openMouthImg   = '/static/images/char-mouth-open.png';
const closedMouthImg = '/static/images/char-mouth-closed.png';
```

A animação do personagem é feita alternando entre duas imagens estáticas: boca aberta e boca fechada. Simples e eficaz — não precisa de canvas ou WebGL.

---

### `populateVoiceList` — Síntese de voz

```js
const allVoices = speechSynthesis.getVoices();
voices = allVoices.filter(voice => voice.name.includes('Google'));
```

Usa a Web Speech API nativa do browser. Filtra apenas vozes do Google, que são mais naturais. O `select` no HTML permite o usuário escolher a voz antes de enviar.

```js
if (speechSynthesis.onvoiceschanged !== undefined) {
    speechSynthesis.onvoiceschanged = populateVoiceList;
}
```

As vozes podem não estar disponíveis imediatamente no carregamento — o evento `onvoiceschanged` garante que a lista seja preenchida quando o browser terminar de carregar as vozes disponíveis.

---

### `typewriter` — Efeito de digitação

```js
const segmenter = new Intl.Segmenter(undefined, { granularity: 'grapheme' });
const segments = Array.from(segmenter.segment(text)).map(s => s.segment);
```

Em vez de iterar caractere por caractere com `charAt`, usa `Intl.Segmenter` para dividir o texto em grafemas. Isso garante que emojis e caracteres Unicode compostos (ex: 👨‍👩‍👧) apareçam corretamente em vez de quebrar no meio.

```js
function type() {
    if (i < segments.length) {
        element.innerHTML += segments[i];
        i++;
        setTimeout(type, speed);
    }
}
```

Recursão com `setTimeout` em vez de `setInterval` — garante que o próximo caractere só apareça depois do anterior ser renderizado, evitando acúmulo de chamadas se o browser estiver lento.

---

### `speak` — Voz + Lip Sync

```js
utterance.onstart = () => {
    let mouthOpen = true;
    lipSyncInterval = setInterval(() => {
        characterImage.src = mouthOpen ? openMouthImg : closedMouthImg;
        mouthOpen = !mouthOpen;
    }, 150);
};
```

Quando a síntese de voz começa, um `setInterval` alterna a imagem do personagem a cada 150ms — criando a ilusão de lip sync. É uma aproximação simples mas visualmente convincente.

```js
utterance.onend = () => {
    clearInterval(lipSyncInterval);
    characterImage.src = closedMouthImg;
};
```

Quando termina de falar, para o intervalo e fecha a boca do personagem.

---

### `handleSendMessage` — Envio da mensagem

```js
const response = await fetch('/chat', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ message: message }),
});
```

POST para o endpoint `/chat` do servidor Go com o JSON `{ "message": "..." }`. A resposta é um JSON `{ "response": "..." }`.

```js
typewriter(data.response, status);
speak(data.response);
```

As duas ações acontecem ao mesmo tempo: o texto aparece com efeito de digitação enquanto a voz sintetizada fala — e o personagem anima a boca sincronizado com a fala.

---

## `save_credentials.sh` — Salvar Credenciais

```bash
read -p "Please enter your Google Cloud project ID: " user_project_id
echo "$user_project_id" > "$PROJECT_FILE"
```

Script interativo que pede duas informações ao usuário:
1. **Project ID** do Google Cloud — necessário para configurar o `gcloud`
2. **API Key** do AI Studio — usada pelo ADK e pelo nano-banana para chamar a API Gemini

As credenciais são salvas em arquivos de texto simples no home directory:
- `~/project_id.txt`
- `~/api_key.txt`

**Por que salvar em arquivo?** O Cloud Shell reinicia periodicamente e perde variáveis de ambiente. Salvar em arquivo garante que as credenciais persistam entre sessões — o `setup.sh` lê esses arquivos toda vez que é executado.

**Por que `bash ./save_credentials.sh` (não `source`)?** Este script não exporta variáveis — só salva em disco. Pode ser executado como subshell normalmente.

---

## `setup.sh` — Configurar Ambiente

```bash
# IMPORTANT: Must be SOURCED so exports persist in your shell.
# Usage: source ./setup.sh
```

Diferente do `save_credentials.sh`, este **precisa** ser executado com `source` (ou `. ./setup.sh`). Sem `source`, os `export` acontecem num subshell e as variáveis somem quando o script termina.

---

### Verificação de arquivos

```bash
if [[ ! -f "$PROJECT_FILE" ]]; then
    handle_error "Project ID file not found. Run: bash ./save_credentials.sh"
    return 1
fi
```

Usa `return 1` (não `exit 1`) para não fechar o terminal quando chamado com `source`. `exit` encerraria o shell inteiro do usuário.

---

### Verificação de autenticação

```bash
if gcloud auth print-access-token > /dev/null 2>&1; then
    echo "gcloud is authenticated."
else
    ...
    return 1
fi
```

Testa se o `gcloud` está autenticado tentando gerar um token de acesso. `> /dev/null 2>&1` descarta tanto stdout quanto stderr — só o código de saída importa.

---

### Exports

```bash
export GOOGLE_API_KEY="$user_api_key"
export PROJECT_ID=$(gcloud config get project)
```

`GOOGLE_API_KEY` é lida pelo ADK, pelo nano-banana e pelo Gemini CLI — é a variável mais importante do setup. `PROJECT_ID` é útil para comandos `gcloud` durante o workshop.

---

### Persistência automática no `~/.bashrc`

```bash
echo 'export GOOGLE_API_KEY=$(cat ~/api_key.txt)' >> ~/.bashrc
```

Esta linha é adicionada manualmente pelo participante no Step 2. Garante que qualquer novo terminal (como o que vai rodar o nano-banana) já tenha a variável disponível sem precisar executar `source ./setup.sh` novamente.

---

## `nano-banana-mcp/main.go` — Servidor MCP de Geração de Imagens

Servidor MCP independente que expõe ferramentas de geração de imagem via protocolo HTTP + SSE. Roda como processo separado da aplicação principal.

---

### Constante do modelo

```go
const imageModel = "gemini-2.5-flash-image"
```

Modelo específico para geração de imagens via AI Studio. Diferente dos modelos de texto, ele retorna dados binários (PNG) dentro da resposta.

---

### `newGeminiClient`

```go
func newGeminiClient(ctx context.Context) (*genai.Client, error) {
    if os.Getenv("GOOGLE_API_KEY") == "" {
        return nil, fmt.Errorf("GOOGLE_API_KEY environment variable not set")
    }
    return genai.NewClient(ctx, nil)
}
```

O novo SDK `google.golang.org/genai` lê `GOOGLE_API_KEY` do ambiente automaticamente quando `nil` é passado como config — não precisa de configuração explícita. A verificação manual antes serve para dar uma mensagem de erro clara em vez de falhar silenciosamente dentro do SDK.

---

### `extractImageData`

```go
for _, part := range candidate.Content.Parts {
    if part.InlineData != nil {
        return part.InlineData.Data, nil
    }
}
```

A resposta de geração de imagem contém múltiplos `Parts` — texto e dados binários misturados. O loop procura o primeiro part com `InlineData` (dados da imagem em bytes) e ignora os parts de texto.

---

### `generateLipSyncImagesHandler` — Geração em dois passos

```go
// Passo 1: gera boca aberta
respOpen, err := client.Models.GenerateContent(ctx, imageModel,
    genai.Text(params.Arguments.Prompt+" with mouth open"), imageConfig)

// Passo 2: usa a imagem anterior como referência
closedContents := []*genai.Content{{
    Role: genai.RoleUser,
    Parts: []*genai.Part{
        genai.NewPartFromText("change the mouth from open to close"),
        genai.NewPartFromBytes(openData, "image/png"),
    },
}}
respClosed, err := client.Models.GenerateContent(ctx, imageModel, closedContents, imageConfig)
```

A técnica de consistência visual: em vez de gerar as duas imagens de forma independente (o que resultaria em personagens diferentes), a imagem com boca fechada é gerada a partir da imagem com boca aberta como referência. O modelo mantém o mesmo personagem e só altera a posição da boca.

---

### `imageConfig` — Modalidade de resposta

```go
var imageConfig = &genai.GenerateContentConfig{
    ResponseModalities: []string{"IMAGE", "TEXT"},
}
```

Instrui o modelo a retornar dados de imagem na resposta. Sem essa configuração, o modelo de geração de imagens devolveria apenas texto descrevendo a imagem — não os bytes da imagem em si.

---

### `textResult` — Retorno para o MCP

```go
func textResult(data any) (*mcp.CallToolResultFor[any], error) {
    b, _ := json.Marshal(data)
    return &mcp.CallToolResultFor[any]{
        Content:           []mcp.Content{&mcp.TextContent{Text: string(b)}},
        StructuredContent: data,
    }, nil
}
```

O protocolo MCP retorna resultados como texto ou conteúdo estruturado. As ferramentas de imagem retornam os **caminhos** dos arquivos salvos (não os bytes da imagem) — o Gemini CLI recebe esses caminhos e pode então mover ou referenciar os arquivos.

---

### `main` — Registro das ferramentas e servidor SSE

```go
s := mcp.NewServer(&mcp.Implementation{Name: "nano_banana", Version: "1.0.0"}, nil)

mcp.AddTool(s, &mcp.Tool{
    Name:        "generate_lip_sync_images",
    Description: "Generates two images for a lip-syncing app, one with mouth open and one with mouth closed.",
}, generateLipSyncImagesHandler)
```

A `Description` de cada tool é enviada ao LLM junto com a definição da função. É o texto que o Gemini lê para decidir quando e como chamar a tool — deve ser clara e específica.

```go
handler := mcp.NewSSEHandler(func(r *http.Request) *mcp.Server { return s })
http.Handle("/", handler)
http.ListenAndServe(":8080", nil)
```

SSE (Server-Sent Events) é o transporte: o cliente abre uma conexão HTTP persistente e o servidor envia eventos de texto conforme o processamento avança. É mais simples que WebSocket e suficiente para o padrão de comunicação do MCP.
