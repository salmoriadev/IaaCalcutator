# 🔧 Installation & Setup

## ✅ Ubuntu (apt) or Arch (recommended)

1. **Install Go** (1.21+). Download: https://golang.org/dl/

2. **Install Fyne dependencies (CGO + OpenGL + X11)**.
Ubuntu/Debian (apt):
```bash
sudo apt update
sudo apt install -y gcc libgl1-mesa-dev xorg-dev
```
Arch:
```bash
sudo pacman -S --needed gcc mesa libx11 libxcursor libxrandr libxinerama libxi
```

3. **Install Go dependencies**:
```bash
go mod tidy
```

4. **Run**:
```bash
go run .
```

## ✅ Windows (if needed)

1. **Install a C compiler**.
Option A: TDM-GCC (64-bit) - https://jmeubank.github.io/tdm-gcc/
Option B: MSYS2 - https://www.msys2.org/ and `pacman -S mingw-w64-x86_64-gcc`

2. **Add the compiler to PATH**:
TDM-GCC: `C:\TDM-GCC-64\bin`
MSYS2: `C:\msys64\mingw64\bin`

3. **Enable CGO** and run `go run .`.

## ✅ Alternative (software renderer)

If OpenGL is not available, try:
```bash
go build -tags software -o iaa-calculator
```

## 📝 Note

If you hit dependency errors on Linux, recheck the X11/OpenGL packages above. On Windows, the most common error is "gcc not found".
