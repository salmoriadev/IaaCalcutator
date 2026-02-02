# Script de build para Windows PowerShell
# Habilita CGO e compila a aplicação

Write-Host "🔧 Habilitando CGO..." -ForegroundColor Yellow
$env:CGO_ENABLED = "1"

Write-Host "📦 Verificando dependências..." -ForegroundColor Yellow
go mod tidy

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Erro ao verificar dependências" -ForegroundColor Red
    exit 1
}

Write-Host "🔨 Compilando aplicação..." -ForegroundColor Yellow
go build -o calculadora-iaa.exe

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Compilação concluída com sucesso!" -ForegroundColor Green
    Write-Host "📁 Executável criado: calculadora-iaa.exe" -ForegroundColor Green
} else {
    Write-Host "❌ Erro na compilação" -ForegroundColor Red
    Write-Host "💡 Verifique se o compilador C (gcc) está instalado e no PATH" -ForegroundColor Yellow
    Write-Host "📖 Veja INSTALACAO.md para mais detalhes" -ForegroundColor Yellow
    exit 1
}

