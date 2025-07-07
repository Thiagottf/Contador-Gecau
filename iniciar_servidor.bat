@echo off
echo ================================
echo  Sistema Contador - Banpara
echo ================================
echo Iniciando servidor Go...
echo.

REM Verifica se Go está instalado
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo ERRO: Go não encontrado! Instale Go em https://golang.org/dl/
    pause
    exit /b 1
)

REM Navega para a pasta raiz do projeto
cd /d "%~dp0"

REM Instala dependências
echo Verificando dependências...
go mod tidy

REM Executa o servidor
echo.
echo Iniciando servidor na porta 8080...
echo Acesse: http://localhost:8080
echo.
echo Pressione Ctrl+C para parar o servidor
echo.
go run main.go

pause
