# 🔧 Instalação e Configuração

## ⚠️ Problema: CGO não configurado

O Fyne precisa do CGO habilitado e de um compilador C no Windows.

## ✅ Solução 1: Instalar TDM-GCC (Recomendado)

1. **Baixe o TDM-GCC**:
   - https://jmeubank.github.io/tdm-gcc/
   - Escolha a versão 64-bit
   - Instale com as opções padrão

2. **Configure o PATH**:
   - Adicione `C:\TDM-GCC-64\bin` ao PATH do sistema
   - Reinicie o terminal/PowerShell

3. **Habilite o CGO**:
   ```powershell
   $env:CGO_ENABLED=1
   ```

4. **Teste**:
   ```bash
   go run .
   ```

## ✅ Solução 2: Usar MSYS2

1. **Instale o MSYS2**:
   - https://www.msys2.org/
   - Execute: `pacman -S mingw-w64-x86_64-gcc`
   - Adicione ao PATH: `C:\msys64\mingw64\bin`

2. **Habilite o CGO**:
   ```powershell
   $env:CGO_ENABLED=1
   ```

## ✅ Solução 3: Compilar com tags (Software Renderer)

Tente compilar com o driver software:

```bash
go build -tags software -o calculadora-iaa.exe
```

## 📝 Nota

Se nenhuma das soluções funcionar, considere usar uma versão web-based ou outra biblioteca GUI que não precise de CGO.

