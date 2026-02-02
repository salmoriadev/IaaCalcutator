package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// App represents the GUI application.
type App struct {
	app              fyne.App
	window           fyne.Window
	currentIAA       *widget.Entry
	completedCredits *widget.Entry
	courses          []Course
	coursesTable     *widget.Table
	resultLabel      *widget.Label
}

// NewApp creates a new application instance.
func NewApp() *App {
	myApp := app.NewWithID("iaa.calculator")
	window := myApp.NewWindow("📊 IAA Calculator")
	window.Resize(fyne.NewSize(900, 720))
	window.CenterOnScreen()

	return &App{
		app:     myApp,
		window:  window,
		courses: []Course{},
	}
}

// Start shows the home screen and starts the app loop.
func (a *App) Start() {
	a.createHomeScreen()
	a.window.ShowAndRun()
}
