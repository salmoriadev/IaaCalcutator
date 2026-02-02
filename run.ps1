# Script para executar a aplicação
# Habilita CGO e executa

Write-Host "🔧 Habilitando CGO..." -ForegroundColor Yellow
$env:CGO_ENABLED = "1"

Write-Host "🚀 Iniciando aplicação..." -ForegroundColor Yellow
go run .

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Erro ao executar" -ForegroundColor Red
    Write-Host "💡 Verifique se o compilador C (gcc) está instalado e no PATH" -ForegroundColor Yellow
    Write-Host "📖 Veja INSTALACAO.md para mais detalhes" -ForegroundColor Yellow
}

