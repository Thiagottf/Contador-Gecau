# Script PowerShell para iniciar o Servidor do Sistema Contador Banpara
Write-Host "================================" -ForegroundColor Cyan
Write-Host " Sistema Contador - Banpara" -ForegroundColor Cyan  
Write-Host "================================" -ForegroundColor Cyan
Write-Host "Iniciando servidor Go..." -ForegroundColor Yellow
Write-Host ""

# Verifica se Go está instalado
try {
    $goVersion = go version 2>$null
    Write-Host "Go encontrado: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "ERRO: Go não encontrado!" -ForegroundColor Red
    Write-Host "Instale Go em: https://golang.org/dl/" -ForegroundColor Yellow
    Read-Host "Pressione Enter para sair"
    exit 1
}

# Navega para a pasta raiz do projeto
Set-Location $PSScriptRoot

# Instala dependências
Write-Host "Verificando dependências..." -ForegroundColor Yellow
go mod tidy

# Executa o servidor
Write-Host ""
Write-Host "Iniciando servidor na porta 8080..." -ForegroundColor Green
Write-Host "Acesse: http://localhost:8080" -ForegroundColor Cyan
Write-Host ""
Write-Host "Pressione Ctrl+C para parar o servidor" -ForegroundColor Yellow
Write-Host ""

try {
    go run main.go
}
catch {
    Write-Host "Erro ao iniciar o servidor: $($_.Exception.Message)" -ForegroundColor Red
    Read-Host "Pressione Enter para sair"
}
    Read-Host "Pressione Enter para sair"
}
