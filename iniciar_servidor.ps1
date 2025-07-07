Write-Host "=== Sistema Contador Banpara ===" -ForegroundColor Green
Write-Host "Iniciando servidor..." -ForegroundColor Yellow
Write-Host ""

# Navega para o diretório correto
Set-Location -Path "$PSScriptRoot\syles"

# Informa onde acessar
Write-Host "Servidor rodando na porta 8080" -ForegroundColor Cyan
Write-Host "Acesse: http://localhost:8080" -ForegroundColor Cyan
Write-Host ""
Write-Host "Pressione Ctrl+C para parar o servidor" -ForegroundColor Yellow
Write-Host ""

# Executa o servidor
try {
    go run main.go
}
catch {
    Write-Host "Erro ao iniciar o servidor: $($_.Exception.Message)" -ForegroundColor Red
    Read-Host "Pressione Enter para sair"
}
