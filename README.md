# 📊 Calculadora de IAA em Go

Calculadora de IAA (Índice de Aproveitamento Acadêmico) com interface gráfica desenvolvida em Go usando Fyne.

## 🚀 Funcionalidades

- ✅ **Calcular IAA Atualizado**: Adiciona novas matérias e calcula o novo IAA
- ✅ **Calcular Meta de IAA**: Calcula a média necessária para alcançar um IAA desejado
- ✅ Interface gráfica moderna e intuitiva
- ✅ Gerenciamento de matérias (adicionar, editar, remover)

## 📦 Instalação

### Pré-requisitos

1. **Instale o Go** (versão 1.21 ou superior):
   - Download: https://golang.org/dl/

2. **Instale um compilador C** (necessário para Fyne no Windows):
   - **Opção 1 (Recomendado)**: TDM-GCC - https://jmeubank.github.io/tdm-gcc/
     - Baixe a versão 64-bit
     - Instale com opções padrão
     - Adicione `C:\TDM-GCC-64\bin` ao PATH do sistema
   - **Opção 2**: MSYS2 - https://www.msys2.org/
     - Instale e execute: `pacman -S mingw-w64-x86_64-gcc`
     - Adicione `C:\msys64\mingw64\bin` ao PATH

3. **Instale as dependências**:
   ```bash
   go mod tidy
   ```

### ⚠️ Importante: Habilitar CGO

No PowerShell, antes de executar:
```powershell
$env:CGO_ENABLED=1
```

## 🎯 Execução

### Desenvolvimento

**Opção 1: Usar script (Recomendado)**
```powershell
.\run.ps1
```

**Opção 2: Manual no PowerShell:**
```powershell
$env:CGO_ENABLED=1
go run .
```

**Opção 3: Manual no CMD:**
```cmd
set CGO_ENABLED=1
go run .
```

### Compilar executável

**Opção 1: Usar script (Recomendado)**
```powershell
.\build.ps1
.\calculadora-iaa.exe
```

**Opção 2: Manual no PowerShell:**
```powershell
$env:CGO_ENABLED=1
go build -o calculadora-iaa.exe
.\calculadora-iaa.exe
```

**Opção 3: Manual no CMD:**
```cmd
set CGO_ENABLED=1
go build -o calculadora-iaa.exe
calculadora-iaa.exe
```

### 🔧 Solução de Problemas

Se aparecer erro sobre "gcc not found":
1. Verifique se o compilador C está instalado
2. Verifique se está no PATH do sistema
3. Reinicie o terminal após adicionar ao PATH
4. Veja mais detalhes em `INSTALACAO.md`

## 📁 Estrutura do Projeto

```
IaaCalculator/
├── main.go          # Ponto de entrada da aplicação
├── models.go        # Modelos de dados (Materia)
├── calculos.go      # Lógica de cálculos de IAA
├── gui.go           # Interface gráfica
├── go.mod           # Dependências
└── README.md        # Este arquivo
```

## 🎨 Interface

A aplicação possui uma interface gráfica moderna com:
- Tela inicial com menu de opções
- Formulários organizados em cards
- Lista de matérias com opções de editar e remover
- Mensagens de resultado claras e coloridas

## 📝 Como Usar

1. **Calcular IAA Atualizado**:
   - Informe seu IAA atual e créditos já cursados
   - Adicione as matérias com seus créditos e notas
   - Clique em "Calcular IAA" para ver o resultado

2. **Calcular Meta de IAA**:
   - Informe seu IAA atual e créditos já cursados
   - Informe os créditos do semestre atual
   - Informe o IAA que deseja alcançar
   - Clique em "Calcular Meta" para ver a média necessária

## 🔧 Tecnologias

- **Go 1.21+**: Linguagem de programação
- **Fyne v2**: Framework para interface gráfica

## 📄 Licença

Este projeto é de uso livre para fins educacionais.

