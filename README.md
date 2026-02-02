# 📊 IAA Calculator in Go

IAA (Academic Performance Index) calculator with a graphical interface built in Go using Fyne.

## 🚀 Features

- ✅ **Calculate Updated IAA**: Add new courses and calculate the updated IAA
- ✅ **Calculate IAA Target**: Compute the required average to reach a target IAA
- ✅ Modern, intuitive GUI
- ✅ Course management (add, edit, remove)

## 📦 Installation

### Prerequisites

1. **Install Go** (version 1.21 or higher). Download: https://golang.org/dl/

2. **Linux: install Fyne dependencies (CGO + OpenGL + X11)**.
Ubuntu/Debian (apt):
```bash
sudo apt update
sudo apt install -y gcc libgl1-mesa-dev xorg-dev
```
Arch:
```bash
sudo pacman -S --needed gcc mesa libx11 libxcursor libxrandr libxinerama libxi
```

3. **Windows (if needed)**: install a C compiler (TDM-GCC or MSYS2) and add it to PATH.

4. **Install Go dependencies**:
```bash
go mod tidy
```

### ⚠️ Important: CGO

- **Linux**: usually already enabled (`go env CGO_ENABLED` should return `1`).
- **Windows (PowerShell)**:
```powershell
$env:CGO_ENABLED=1
```

## 🎯 Running

### Development

**Ubuntu/Debian or Arch (recommended)**:
```bash
go run .
```

**Windows (manual)**: enable `CGO_ENABLED=1` and run `go run .`.

### Build executable

**Ubuntu/Debian or Arch**:
```bash
go build -o iaa-calculator
./iaa-calculator
```

**Windows (manual)**: enable `CGO_ENABLED=1` and run `go build -o iaa-calculator.exe`.

### 🔧 Troubleshooting

If you see Linux dependency errors (X11/OpenGL), install the packages listed under **Prerequisites**.
On Windows, if you see "gcc not found", make sure a C compiler is installed and on PATH.
See more details in `INSTALACAO.md`.

## 📁 Project Structure

```
IaaCalculator/
├── main.go          # Application entry point
├── models.go        # Data models (Course)
├── calculos.go      # IAA calculation logic
├── app.go           # App and initialization
├── ui_helpers.go    # UI helpers and validation
├── ui_home.go       # Home screen
├── ui_iaa.go        # Updated IAA screen
├── ui_meta.go       # IAA target screen
├── go.mod           # Dependencies
└── README.md        # This file
```

## 🎨 Interface

The app includes:
- Home screen with menu options
- Card-based forms
- Course table with edit/remove actions
- Clear, readable result messages

## 📝 How to Use

1. **Calculate Updated IAA**:
   - Enter current IAA and completed credits
   - Add courses with credits and grades
   - Click "Calculate IAA" to see the result

2. **Calculate IAA Target**:
   - Enter current IAA and completed credits
   - Enter the current semester credits
   - Enter the target IAA
   - Click "Calculate Target" to see the required average

## 🔧 Tech

- **Go 1.21+**
- **Fyne v2**

## 📄 License

This project is free to use for educational purposes.
