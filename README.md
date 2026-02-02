# IAA Calculator (UFSC GPA Equivalent)

IAA (Índice de Aproveitamento Acadêmico) is the UFSC (Universidade Federal de Santa Catarina) equivalent of GPA. This project provides a clean, desktop GUI to calculate updated IAA and estimate the target average needed to reach a desired IAA.

## Overview

The application is built with Go and Fyne and focuses on clarity and usability:
- Calculate updated IAA after adding new courses
- Calculate required average to reach a target IAA
- Manage courses (add, edit, remove)

## How It Works

The calculator combines your current IAA and completed credits with the grades and credits of new courses. For target calculations, it estimates the average you need in a new semester to reach a desired IAA.

## Requirements

- Go 1.21+
- CGO enabled
- System dependencies for Fyne (OpenGL + X11 on Linux)

### Linux Dependencies

Ubuntu/Debian (apt):
```bash
sudo apt update
sudo apt install -y gcc libgl1-mesa-dev xorg-dev
```

Arch:
```bash
sudo pacman -S --needed gcc mesa libx11 libxcursor libxrandr libxinerama libxi
```

### Windows (if needed)

Install a C compiler (TDM-GCC or MSYS2) and add it to PATH, then enable CGO:
```powershell
$env:CGO_ENABLED=1
```

## Running

```bash
go mod tidy
go run .
```

## Build

Linux:
```bash
go build -o iaa-calculator
./iaa-calculator
```

Windows:
```powershell
$env:CGO_ENABLED=1
go build -o iaa-calculator.exe
```

## Project Structure

```
IaaCalculator/
├── app.go           # App and initialization
├── calculos.go      # Calculation logic
├── main.go          # Entry point
├── models.go        # Data models
├── ui_helpers.go    # UI helpers and validation
├── ui_home.go       # Home screen
├── ui_iaa.go        # Updated IAA screen
├── ui_meta.go       # IAA target screen
└── README.md        # Documentation
```

## Notes

If you encounter dependency errors on Linux, recheck the X11/OpenGL packages above. On Windows, the most common issue is "gcc not found".
