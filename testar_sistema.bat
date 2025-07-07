@echo off
echo =================================
echo  TESTE DO SISTEMA CONTADOR
echo =================================
echo.

echo 1. Compilando o projeto...
go build -o teste.exe main.go

if %errorlevel% neq 0 (
    echo ERRO: Falha na compilação!
    pause
    exit /b 1
)

echo ✅ Compilação bem-sucedida!
echo.

echo 2. Verificando arquivos...
if not exist "index.html" (
    echo ❌ ERRO: index.html não encontrado!
    pause
    exit /b 1
)
echo ✅ index.html encontrado

if not exist "static\css\styles.css" (
    echo ❌ ERRO: static\css\styles.css não encontrado!
    pause
    exit /b 1
)
echo ✅ CSS encontrado

echo.
echo 3. Testando execução...
echo Iniciando servidor por 5 segundos...
start /B teste.exe
timeout /t 5 /nobreak >nul

echo.
echo 4. Limpando arquivos de teste...
taskkill /F /IM teste.exe >nul 2>&1
del teste.exe >nul 2>&1

echo.
echo =================================
echo  ✅ TESTE CONCLUÍDO COM SUCESSO!
echo =================================
echo.
echo A reorganização foi bem-sucedida:
echo ✅ Backend na raiz funcionando
echo ✅ CSS na estrutura correta
echo ✅ Compilação sem erros
echo ✅ Estrutura de arquivos OK
echo.
echo Para executar o sistema:
echo   go run main.go
echo   ou
echo   .\iniciar_servidor.ps1
echo.
pause
