package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createCard(title string, content fyne.CanvasObject) *widget.Card {
	return widget.NewCard(title, "", content)
}

func (a *App) fieldWithLabel(labelText string, entry *widget.Entry) fyne.CanvasObject {
	return container.NewBorder(nil, nil, widget.NewLabel(labelText), nil, entry)
}

func (a *App) validateBaseInputs() (float64, int, error) {
	currentIAAStr := strings.TrimSpace(a.currentIAA.Text)
	currentIAA, err := strconv.ParseFloat(currentIAAStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid IAA")
	}

	if currentIAA < 0 || currentIAA > 10 {
		return 0, 0, fmt.Errorf("IAA must be between 0 and 10")
	}

	completedCreditsStr := strings.TrimSpace(a.completedCredits.Text)
	completedCredits, err := strconv.Atoi(completedCreditsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid credits")
	}

	if completedCredits < 0 {
		return 0, 0, fmt.Errorf("credits must be non-negative")
	}

	return currentIAA, completedCredits, nil
}

func allFilled(entries ...*widget.Entry) bool {
	for _, entry := range entries {
		if strings.TrimSpace(entry.Text) == "" {
			return false
		}
	}
	return true
}
