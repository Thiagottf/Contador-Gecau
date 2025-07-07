# Sistema Contador Banpara

Um sistema web moderno para acompanhar quantos dias o sistema está funcionando sem falhas, com sistema de apostas integrado.

## Funcionalidades

### 🔢 Contador Principal
- Exibe quantos dias o sistema está funcionando consecutivamente
- Mostra data de início e último reset
- Indicador visual do status do sistema (ativo/inativo)
- Atualização em tempo real a cada 5 segundos

### ⚙️ Controles do Sistema
- **Definir Data de Início**: Configure quando o sistema começou a funcionar
- **Sistema Caiu**: Marca o sistema como inativo e para a contagem
- **Reiniciar Contador**: Reinicia a contagem do zero

### 🎲 Sistema de Apostas
- Funcionários podem fazer apostas sobre quantos dias o sistema vai aguentar
- Cada aposta inclui:
  - Nome do apostador
  - Número de dias apostados
  - Prêmio da aposta (dinheiro, almoço, etc.)
- Status automático das apostas (Ativa, Ganhou, Perdeu)
- Remoção de apostas

## Tecnologias Utilizadas

### Backend
- **Go (Golang)** - Linguagem principal
- **Gorilla Mux** - Roteamento HTTP
- **CORS** - Suporte a requisições cross-origin

### Frontend
- **HTML5** - Estrutura
- **CSS3** - Estilização moderna com gradientes e animações
- **JavaScript (Vanilla)** - Interatividade
- **Font Awesome** - Ícones
- **Google Fonts (Inter)** - Tipografia

## Instalação e Configuração

### Pré-requisitos
- Go 1.21 ou superior instalado
- Navegador web moderno

### Passo a Passo

1. **Clone ou baixe o projeto**
   ```
   Certifique-se de ter todos os arquivos na estrutura correta
   ```

2. **Instale as dependências do Go**
   ```powershell
   cd ProjetoCotador
   go mod tidy
   ```

3. **Execute o servidor**
   
   **Opção A - Usando Go diretamente:**
   ```powershell
   go run main.go
   ```
   
   **Opção B - Usando scripts automatizados:**
   ```powershell
   # Windows (PowerShell)
   .\iniciar_servidor.ps1
   
   # Windows (Prompt de Comando)
   iniciar_servidor.bat
   ```

4. **Acesse o sistema**
   ```
   Abra seu navegador e vá para: http://localhost:8080
   ```

## Estrutura do Projeto

```
ProjetoCotador/
├── main.go                    # Servidor backend em Go (PRINCIPAL)
├── index.html                 # Página principal do frontend
├── go.mod                     # Dependências do Go
├── go.sum                     # Checksums das dependências
├── iniciar_servidor.bat       # Script de execução (Windows CMD)
├── iniciar_servidor.ps1       # Script de execução (PowerShell)
├── .gitignore                 # Arquivos ignorados pelo Git
├── static/                    # Arquivos estáticos do frontend
│   ├── css/
│   │   └── styles.css         # Estilos CSS principais
│   └── js/                    # Scripts JavaScript (futuro)
├── DOCUMENTACAO.md            # Documentação técnica detalhada
└── README.md                  # Este arquivo
```

## API Endpoints

### Sistema
- `GET /api/status` - Obtém status atual do sistema
- `POST /api/start-date` - Define data de início
- `POST /api/reset` - Marca sistema como caído
- `POST /api/restart` - Reinicia contador

### Apostas
- `GET /api/bets` - Lista todas as apostas
- `POST /api/bets` - Cria nova aposta
- `DELETE /api/bets/{id}` - Remove aposta

## Exemplos de Uso

### Definir Data de Início
```json
POST /api/start-date
{
  "start_date": "2025-01-01"
}
```

### Criar Aposta
```json
POST /api/bets
{
  "name": "João Silva",
  "bet_days": 30,
  "bet_item": "Almoço no restaurante"
}
```

## Funcionalidades da Interface

### 🎨 Design Moderno
- Interface responsiva que funciona em desktop e mobile
- Gradientes e efeitos visuais modernos
- Animações suaves e feedback visual
- Tema profissional adequado para ambiente corporativo

### 📱 Responsividade
- Layout adaptativo para diferentes tamanhos de tela
- Otimizado para uso em celulares e tablets
- Navegação intuitiva em qualquer dispositivo

### 🔔 Notificações
- Sistema de toast para feedback das ações
- Modais de confirmação para ações críticas
- Indicadores visuais de status

### ⚡ Tempo Real
- Atualização automática do contador
- Status das apostas em tempo real
- Sincronização automática com o servidor

## Personalização

### Cores e Tema
O arquivo `static/css/styles.css` pode ser editado para personalizar:
- Cores principais (atualmente azul/roxo)
- Fontes e tipografia
- Espaçamentos e layouts
- Animações e efeitos

### Funcionalidades Adicionais
O código está estruturado para fácil extensão:
- Adicionar novos tipos de apostas
- Implementar sistema de usuários
- Adicionar histórico de resets
- Integrar com banco de dados

## Possíveis Melhorias Futuras

1. **Persistência de Dados**
   - Integração com banco de dados (PostgreSQL, MySQL)
   - Backup automático dos dados

2. **Sistema de Usuários**
   - Login e autenticação
   - Perfis de usuários
   - Controle de permissões

3. **Relatórios e Estatísticas**
   - Gráficos de uptime
   - Histórico de falhas
   - Estatísticas das apostas

4. **Notificações**
   - Email quando sistema cai
   - Alertas de apostas vencidas
   - Webhooks para integração

## Suporte

Para dúvidas ou problemas:
1. Verifique se todas as dependências estão instaladas
2. Confirme que a porta 8080 está disponível
3. Verifique os logs do servidor no terminal

## Licença

Este projeto foi desenvolvido para uso interno do Banpara.
