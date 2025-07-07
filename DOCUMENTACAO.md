# Sistema de Contador de Dias - Documentação Técnica

## Visão Geral

Este sistema foi desenvolvido para monitorar quantos dias consecutivos um serviço/sistema funciona sem falhas, além de permitir que usuários façam apostas sobre quando o sistema poderá falhar.

## Arquitetura do Projeto

### Backend (Go)
- **Linguagem**: Go (Golang)
- **Framework**: Gorilla Mux (roteamento HTTP)
- **Porta**: 8080
- **CORS**: Habilitado para permitir requisições do frontend

### Frontend 
- **HTML5**: Estrutura da página
- **CSS3**: Estilização moderna com design responsivo
- **JavaScript**: Interação dinâmica com o backend
- **Fontes**: Google Fonts (Inter)
- **Ícones**: Font Awesome 6.0

## Estrutura de Arquivos

```
ProjetoCotador/
├── syles/
│   ├── main.go          # Backend em Go
│   ├── main.exe         # Executável compilado
│   └── syles.css        # Estilos CSS
├── index.html           # Frontend da aplicação
├── go.mod              # Dependências do Go
├── go.sum              # Checksums das dependências
├── README.md           # Instruções básicas
├── DOCUMENTACAO.md     # Esta documentação
├── iniciar_servidor.bat # Script Windows (Batch)
└── iniciar_servidor.ps1 # Script PowerShell
```

## Backend - Explicação Detalhada

### Estruturas de Dados

#### SystemStatus
```go
type SystemStatus struct {
    StartDate    time.Time `json:"start_date"`    // Data de início do sistema
    DaysRunning  int       `json:"days_running"`  // Dias consecutivos rodando
    LastReset    time.Time `json:"last_reset"`    // Data do último reset
    IsRunning    bool      `json:"is_running"`    // Status atual (ativo/inativo)
}
```

#### Bet (Aposta)
```go
type Bet struct {
    ID       int       `json:"id"`        // ID único da aposta
    Name     string    `json:"name"`      // Nome do apostador
    BetDays  int       `json:"bet_days"`  // Quantos dias apostou
    BetItem  string    `json:"bet_item"`  // Prêmio da aposta
    Date     time.Time `json:"date"`      // Data da aposta
    Status   string    `json:"status"`    // "active", "won", "lost"
}
```

### API Endpoints

#### GET /api/status
- **Função**: Retorna o status atual do sistema
- **Resposta**: JSON com SystemStatus
- **Uso**: Atualizar o contador na interface

#### POST /api/start-date
- **Função**: Define uma nova data de início
- **Body**: `{"start_date": "2023-01-01"}`
- **Formato**: YYYY-MM-DD
- **Uso**: Configurar início personalizado

#### POST /api/reset
- **Função**: Marca o sistema como "caído"
- **Efeito**: Para a contagem, mantém o valor atual
- **Uso**: Quando o sistema falha

#### POST /api/restart
- **Função**: Reinicia completamente o contador
- **Efeito**: Zera tudo e recomeça
- **Uso**: Novo ciclo de monitoramento

#### GET /api/bets
- **Função**: Lista todas as apostas
- **Resposta**: Array JSON de apostas
- **Uso**: Exibir apostas na interface

#### POST /api/bets
- **Função**: Cria uma nova aposta
- **Body**: `{"name": "João", "bet_days": 30, "bet_item": "R$ 50"}`
- **Uso**: Fazer nova aposta

#### DELETE /api/bets/{id}
- **Função**: Remove uma aposta específica
- **Parâmetro**: ID da aposta na URL
- **Uso**: Excluir aposta

### Funções Principais

#### calculateDaysRunning()
- Calcula dias decorridos desde o início
- Se sistema parado: mantém último valor
- Se ativo: calcula diferença até agora

#### updateBetsStatus()
- Atualiza status de todas as apostas
- Verifica se alguém ganhou/perdeu
- Chamada após mudanças no sistema

## Frontend - Explicação Detalhada

### JavaScript - Funções Principais

#### loadSystemStatus()
- Busca status do servidor a cada 5 segundos
- Atualiza contador visual
- Atualiza indicador de status (verde/vermelho)

#### loadBets()
- Carrega lista de apostas do servidor
- Atualiza display de apostas

#### displayBets(bets)
- Renderiza apostas na interface
- Aplica cores baseadas no status
- Gera HTML dinamicamente

#### setStartDate()
- Envia nova data de início para o servidor
- Valida entrada do usuário
- Mostra feedback via toast

#### resetSystem() / restartSystem()
- Mostra modal de confirmação
- Executa ação após confirmação
- Atualiza interface

#### createBet()
- Valida dados do formulário
- Envia aposta para servidor
- Limpa formulário após sucesso

### CSS - Conceitos de Design

#### Sistema de Cores
- **Primária**: Azul (#667eea) a Roxo (#764ba2)
- **Sucesso**: Verde (#48bb78)
- **Erro**: Vermelho (#f56565)
- **Aviso**: Laranja (#ed8936)

#### Efeitos Visuais
- **Backdrop Filter**: Efeito vidro fosco
- **Box Shadow**: Sombras para profundidade
- **Transform**: Animações de hover
- **Gradients**: Fundos com gradiente

#### Responsividade
- **Desktop**: Layout de 2-3 colunas
- **Tablet**: Layout ajustado
- **Mobile**: Layout de coluna única

## Como Funciona o Sistema

### 1. Inicialização
- Sistema inicia com data atual
- Contador em 0 dias
- Status "ativo"

### 2. Contagem
- Calcula diferença entre agora e data de início
- Atualiza a cada 5 segundos
- Mostra em tempo real

### 3. Apostas
- Usuários apostam quantos dias o sistema durará
- Sistema monitora automaticamente
- Atualiza status (ganhou/perdeu) conforme necessário

### 4. Reset vs Restart
- **Reset**: Sistema "caiu" - para contagem
- **Restart**: Recomeça do zero

## Fluxo de Dados

1. **Frontend** faz requisição para `/api/status`
2. **Backend** calcula dias rodando
3. **Backend** atualiza status das apostas
4. **Backend** retorna JSON com dados
5. **Frontend** atualiza interface
6. Processo se repete a cada 5 segundos

## Instalação e Execução

### Pré-requisitos
- Go 1.16+ instalado
- Navegador web moderno

### Dependências Go
```bash
go mod tidy
```

### Executar o Servidor
```bash
# Opção 1: Go direto
cd syles
go run main.go

# Opção 2: Compilar e executar
cd syles
go build -o main.exe main.go
./main.exe

# Opção 3: Scripts prontos
# Windows Batch:
iniciar_servidor.bat

# PowerShell:
iniciar_servidor.ps1
```

### Acessar a Aplicação
- Abrir navegador em: http://localhost:8080

## Personalização

### Alterar Porta
```go
// No main.go, linha final:
log.Fatal(http.ListenAndServe(":8080", handler))
// Mudar :8080 para porta desejada
```

### Alterar Intervalo de Atualização
```javascript
// No index.html, função DOMContentLoaded:
setInterval(loadSystemStatus, 5000); // 5000ms = 5 segundos
```

### Alterar Cores
```css
/* No syles.css, alterar variáveis de cor: */
.btn-primary {
    background: linear-gradient(135deg, #667eea, #764ba2);
}
```

## Considerações de Segurança

### Produção
- Configurar CORS adequadamente
- Implementar autenticação se necessário
- Usar HTTPS
- Validar todas as entradas

### Limitações Atuais
- Dados em memória (perdem-se ao reiniciar servidor)
- Sem autenticação
- CORS aberto para qualquer origem

## Possíveis Melhorias

### Backend
- Persistência em banco de dados
- Sistema de usuários
- API de histórico
- Logs estruturados

### Frontend
- PWA (Progressive Web App)
- Notificações push
- Gráficos de histórico
- Tema dark/light

### Funcionalidades
- Sistema de pontuação
- Apostas com prazos
- Backup automático
- Dashboard administrativo

## Troubleshooting

### Servidor não inicia
- Verificar se porta 8080 está livre
- Verificar dependências Go instaladas
- Verificar permissões de firewall

### Frontend não carrega dados
- Verificar se servidor está rodando
- Verificar console do navegador para erros
- Verificar URL da API (localhost:8080)

### Apostas não aparecem
- Verificar se há apostas cadastradas
- Verificar console do navegador
- Testar endpoint /api/bets diretamente

## Logs e Debug

### Backend
```go
// Adicionar logs para debug:
log.Printf("Status atual: %+v", systemStatus)
```

### Frontend
```javascript
// Console do navegador mostra:
console.log("Status carregado:", data);
```

## Contato e Suporte

Para dúvidas sobre implementação ou melhorias, revisar este documento ou consultar os comentários no código fonte.
