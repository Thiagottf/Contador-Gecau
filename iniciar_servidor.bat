@echo off
echo Iniciando Servidor do Sistema Contador Banpara...
echo.
cd /d "%~dp0\syles"
echo Servidor rodando na porta 8080
echo Acesse: http://localhost:8080
echo.
echo Pressione Ctrl+C para parar o servidor
echo.
go run main.go
pause
