// Package main - Sistema de contagem de dias sem falhas e apostas
// Este sistema monitora quantos dias um serviço/sistema funciona sem interrupções
// e permite que usuários façam apostas sobre quando o sistema irá falhar
package main

import (
	"encoding/json" // Para serialização/deserialização de dados JSON
	"fmt"          // Para formatação de strings e output
	"log"          // Para logging de erros
	"net/http"     // Para criar o servidor HTTP
	"strconv"      // Para conversão de strings para números
	"time"         // Para manipulação de datas e tempo

	"github.com/gorilla/mux" // Router HTTP mais avançado
	"github.com/rs/cors"     // Para habilitar CORS (Cross-Origin Resource Sharing)
)

// SystemStatus representa o estado atual do sistema que está sendo monitorado
// Contém informações sobre quando o sistema iniciou, quantos dias está rodando, etc.
type SystemStatus struct {
	StartDate    time.Time `json:"start_date"`    // Data de início do sistema
	DaysRunning  int       `json:"days_running"`  // Número de dias que o sistema está ativo
	LastReset    time.Time `json:"last_reset"`    // Data do último reset/falha do sistema
	IsRunning    bool      `json:"is_running"`    // Se o sistema está atualmente ativo ou não
}

// Bet representa uma aposta feita por um usuário
// Cada aposta contém informações sobre quem apostou, quantos dias apostou e o que foi apostado
type Bet struct {
	ID       int       `json:"id"`        // ID único da aposta
	Name     string    `json:"name"`      // Nome da pessoa que fez a aposta
	BetDays  int       `json:"bet_days"`  // Quantos dias a pessoa apostou que o sistema duraria
	BetItem  string    `json:"bet_item"`  // O que foi apostado (prêmio)
	Date     time.Time `json:"date"`      // Data em que a aposta foi feita
	Status   string    `json:"status"`    // Status da aposta: "active", "won", "lost"
}

// Variáveis globais para manter o estado do sistema
var (
	systemStatus SystemStatus // Estado atual do sistema de monitoramento
	bets         []Bet        // Lista de todas as apostas feitas
	nextBetID    = 1          // Próximo ID disponível para uma nova aposta
)

// init() é executada automaticamente quando o programa inicia
// Inicializa o sistema com valores padrão
func init() {
	// Inicializa o sistema com data atual
	systemStatus = SystemStatus{
		StartDate:   time.Now(), // Sistema começa agora
		DaysRunning: 0,          // Zero dias rodando inicialmente
		LastReset:   time.Now(), // Último reset é agora
		IsRunning:   true,       // Sistema está ativo
	}
}

// calculateDaysRunning calcula quantos dias se passaram desde o início do sistema
// Se o sistema não estiver rodando, retorna o último valor calculado
// Se estiver rodando, calcula baseado na diferença entre agora e a data de início
func calculateDaysRunning() int {
	// Se o sistema não estiver rodando, mantém o último valor
	if !systemStatus.IsRunning {
		return systemStatus.DaysRunning
	}
	// Calcula a diferença de tempo entre agora e o início
	duration := time.Since(systemStatus.StartDate)
	// Converte para dias (24 horas = 1 dia)
	return int(duration.Hours() / 24)
}

// updateBetsStatus atualiza o status das apostas baseado no estado atual do sistema
// Verifica se o sistema parou antes da data apostada (perdeu) ou se chegou na data (ganhou)
func updateBetsStatus() {
	currentDays := calculateDaysRunning()
	for i := range bets {
		// Só atualiza apostas que ainda estão ativas
		if bets[i].Status == "active" {
			// Se o sistema parou antes da data apostada, a aposta perdeu
			if !systemStatus.IsRunning && currentDays < bets[i].BetDays {
				bets[i].Status = "lost"
			// Se chegou ou passou da data apostada, a aposta ganhou
			} else if currentDays >= bets[i].BetDays {
				bets[i].Status = "won"
			}
		}
	}
}

// getStatusHandler retorna o status atual do sistema
// Endpoint: GET /api/status
// Atualiza os dias rodando e status das apostas antes de retornar
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	systemStatus.DaysRunning = calculateDaysRunning()
	updateBetsStatus()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systemStatus)
}

// setStartDateHandler permite definir uma data de início personalizada
// Endpoint: POST /api/start-date
// Recebe: {"start_date": "2023-01-01"}
func setStartDateHandler(w http.ResponseWriter, r *http.Request) {
	// Estrutura para receber a data via JSON
	var req struct {
		StartDate string `json:"start_date"`
	}
	
	// Decodifica o JSON recebido
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Converte a string para time.Time (formato YYYY-MM-DD)
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	
	// Atualiza o sistema com a nova data
	systemStatus.StartDate = startDate
	systemStatus.IsRunning = true
	systemStatus.DaysRunning = calculateDaysRunning()
	updateBetsStatus()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systemStatus)
}

// resetSystemHandler marca o sistema como parado (com falha)
// Endpoint: POST /api/reset
// Não reinicia o contador, apenas marca que o sistema parou
func resetSystemHandler(w http.ResponseWriter, r *http.Request) {
	systemStatus.DaysRunning = calculateDaysRunning()
	systemStatus.LastReset = time.Now()
	systemStatus.IsRunning = false // Marca sistema como parado
	updateBetsStatus() // Atualiza apostas (algumas podem ter perdido)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systemStatus)
}

// restartSystemHandler reinicia completamente o sistema
// Endpoint: POST /api/restart
// Zera o contador e marca sistema como ativo novamente
func restartSystemHandler(w http.ResponseWriter, r *http.Request) {
	systemStatus.StartDate = time.Now()   // Nova data de início
	systemStatus.LastReset = time.Now()   // Atualiza data do último reset
	systemStatus.IsRunning = true         // Sistema volta a estar ativo
	systemStatus.DaysRunning = 0          // Zera contador de dias
	updateBetsStatus()                    // Atualiza status das apostas
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systemStatus)
}

// getBetsHandler retorna todas as apostas atuais
// Endpoint: GET /api/bets
// Atualiza o status das apostas antes de retornar
func getBetsHandler(w http.ResponseWriter, r *http.Request) {
	updateBetsStatus()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bets)
}

// createBetHandler cria uma nova aposta
// Endpoint: POST /api/bets
// Recebe: {"name": "João", "bet_days": 30, "bet_item": "R$ 50"}
func createBetHandler(w http.ResponseWriter, r *http.Request) {
	var newBet Bet
	// Decodifica o JSON recebido
	if err := json.NewDecoder(r.Body).Decode(&newBet); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Valida se todos os campos obrigatórios foram preenchidos
	if newBet.Name == "" || newBet.BetDays <= 0 || newBet.BetItem == "" {
		http.Error(w, "Name, bet_days, and bet_item are required", http.StatusBadRequest)
		return
	}
	
	// Configura a nova aposta
	newBet.ID = nextBetID
	nextBetID++
	newBet.Date = time.Now()
	newBet.Status = "active"
	
	// Adiciona à lista de apostas
	bets = append(bets, newBet)
	updateBetsStatus()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBet)
}

// deleteBetHandler remove uma aposta específica
// Endpoint: DELETE /api/bets/{id}
func deleteBetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	betID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bet ID", http.StatusBadRequest)
		return
	}
	
	// Procura e remove a aposta com o ID especificado
	for i, bet := range bets {
		if bet.ID == betID {
			bets = append(bets[:i], bets[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	http.Error(w, "Bet not found", http.StatusNotFound)
}

// main é a função principal que configura e inicia o servidor
func main() {
	// Cria um novo router usando o Gorilla Mux
	r := mux.NewRouter()
	
	// Configura as rotas da API REST
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/status", getStatusHandler).Methods("GET")          // GET: Obter status do sistema
	api.HandleFunc("/start-date", setStartDateHandler).Methods("POST")   // POST: Definir nova data de início
	api.HandleFunc("/reset", resetSystemHandler).Methods("POST")        // POST: Marcar sistema como parado
	api.HandleFunc("/restart", restartSystemHandler).Methods("POST")    // POST: Reiniciar sistema completamente
	api.HandleFunc("/bets", getBetsHandler).Methods("GET")              // GET: Listar todas as apostas
	api.HandleFunc("/bets", createBetHandler).Methods("POST")           // POST: Criar nova aposta
	api.HandleFunc("/bets/{id}", deleteBetHandler).Methods("DELETE")    // DELETE: Remover aposta específica
	
	// Configura o servidor para servir arquivos estáticos (HTML, CSS, JS)
	// Os arquivos ficam na pasta pai (../) em relação ao diretório atual
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../"))).Methods("GET")
	
	// Configura CORS para permitir requisições do frontend
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                             // Permite qualquer origem
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos HTTP permitidos
		AllowedHeaders: []string{"*"},                             // Headers permitidos
	})
	
	// Aplica o middleware CORS ao router
	handler := c.Handler(r)
	
	// Inicia o servidor HTTP na porta 8080
	fmt.Println("Servidor rodando na porta 8080")
	fmt.Println("Acesse: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}