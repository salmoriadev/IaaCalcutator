package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (a *App) showIAATargetScreen() {
	entrySemesterCredits := widget.NewEntry()
	entrySemesterCredits.SetPlaceHolder("Semester credits")
	entryTargetIAA := widget.NewEntry()
	entryTargetIAA.SetPlaceHolder("Target IAA (0-10)")

	resultLabel := widget.NewLabel("")
	resultLabel.Alignment = fyne.TextAlignCenter
	resultLabel.Wrapping = fyne.TextWrapWord

	hintLabel := widget.NewLabel("")
	hintLabel.Alignment = fyne.TextAlignCenter
	hintLabel.TextStyle = fyne.TextStyle{Italic: true}

	btnCalculate := widget.NewButton("🎯 Calculate Target", nil)
	btnClear := widget.NewButton("🗑️ Clear", func() {
		a.clearTargetFields(entrySemesterCredits, entryTargetIAA, resultLabel)
		updateTargetState(btnCalculate, hintLabel, a.currentIAA, a.completedCredits, entrySemesterCredits, entryTargetIAA)
	})
	btnBack := widget.NewButton("⬅️ Back", a.createHomeScreen)

	update := func() {
		updateTargetState(btnCalculate, hintLabel, a.currentIAA, a.completedCredits, entrySemesterCredits, entryTargetIAA)
	}
	btnCalculate.OnTapped = func() {
		a.calculateIAATarget(entrySemesterCredits, entryTargetIAA, resultLabel)
	}
	a.currentIAA.OnChanged = func(string) { update() }
	a.completedCredits.OnChanged = func(string) { update() }
	entrySemesterCredits.OnChanged = func(string) { update() }
	entryTargetIAA.OnChanged = func(string) { update() }
	update()

	formFrame := a.createCard("Data", container.NewVBox(
		a.fieldWithLabel("Current IAA (0-10):", a.currentIAA),
		a.fieldWithLabel("Completed credits:", a.completedCredits),
		a.fieldWithLabel("Semester credits:", entrySemesterCredits),
		a.fieldWithLabel("Target IAA (0-10):", entryTargetIAA),
	))

	title := widget.NewLabelWithStyle("🎯 Calculate IAA Target", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	subtitle := widget.NewLabel("Enter your data and the target IAA.")
	subtitle.Alignment = fyne.TextAlignCenter
	subtitle.Wrapping = fyne.TextWrapWord

	content := container.NewVBox(
		title,
		subtitle,
		formFrame,
		container.NewGridWithColumns(3, btnCalculate, btnClear, btnBack),
		hintLabel,
		resultLabel,
	)

	a.window.SetContent(container.NewVScroll(content))
}

func updateTargetState(btn *widget.Button, hint *widget.Label, entries ...*widget.Entry) {
	if allFilled(entries...) {
		btn.Enable()
		if hint != nil {
			hint.SetText("")
		}
		return
	}

	btn.Disable()
	if hint != nil {
		hint.SetText("Fill in all fields to calculate.")
	}
}

func (a *App) calculateIAATarget(entrySemesterCredits, entryTargetIAA *widget.Entry, resultLabel *widget.Label) {
	currentIAA, completedCredits, err := a.validateBaseInputs()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	semesterCreditsStr := strings.TrimSpace(entrySemesterCredits.Text)
	semesterCredits, err := strconv.Atoi(semesterCreditsStr)
	if err != nil || semesterCredits <= 0 {
		dialog.ShowError(fmt.Errorf("semester credits must be a positive integer"), a.window)
		return
	}

	targetIAAStr := strings.TrimSpace(entryTargetIAA.Text)
	targetIAA, err := strconv.ParseFloat(targetIAAStr, 64)
	if err != nil {
		dialog.ShowError(fmt.Errorf("invalid target IAA"), a.window)
		return
	}

	if targetIAA < 0 || targetIAA > 10 {
		dialog.ShowError(fmt.Errorf("target IAA must be between 0 and 10"), a.window)
		return
	}

	requiredAverage, pointsNeeded, err := CalculateIAATarget(currentIAA, completedCredits, semesterCredits, targetIAA)
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	var result string
	if requiredAverage > 10 {
		result = "❌ It is not possible to reach this IAA with the planned credits."
	} else if requiredAverage < 0 {
		result = fmt.Sprintf("✅ You have already reached IAA %.2f!\n(required average: %.2f)", targetIAA, requiredAverage)
	} else {
		result = fmt.Sprintf("🎯 To reach IAA %.2f:\n\n📊 Required average: %.2f\n📈 Required total: %.2f",
			targetIAA, requiredAverage, pointsNeeded)
	}

	resultLabel.SetText(result)
}

func (a *App) clearTargetFields(entrySemesterCredits, entryTargetIAA *widget.Entry, resultLabel *widget.Label) {
	a.currentIAA.SetText("")
	a.completedCredits.SetText("")
	entrySemesterCredits.SetText("")
	entryTargetIAA.SetText("")
	resultLabel.SetText("")
}
