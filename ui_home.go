package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// createHomeScreen creates the home screen with the menu.
func (a *App) createHomeScreen() {
	a.currentIAA = widget.NewEntry()
	a.currentIAA.SetPlaceHolder("e.g. 8.5")
	a.completedCredits = widget.NewEntry()
	a.completedCredits.SetPlaceHolder("e.g. 120")

	btnUpdatedIAA := widget.NewButton("📈 Calculate Updated IAA", a.showUpdatedIAAScreen)
	btnTargetIAA := widget.NewButton("🎯 Calculate IAA Target", a.showIAATargetScreen)
	btnExit := widget.NewButton("🚪 Exit", a.window.Close)

	subtitle := widget.NewLabel("Choose an option to continue.")
	subtitle.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		widget.NewLabelWithStyle("📊 IAA Calculator", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		subtitle,
		widget.NewSeparator(),
		btnUpdatedIAA,
		btnTargetIAA,
		widget.NewSeparator(),
		btnExit,
	)

	card := widget.NewCard("", "", content)
	a.window.SetContent(container.NewCenter(card))
}
