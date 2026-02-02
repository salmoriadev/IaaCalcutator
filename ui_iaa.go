package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (a *App) showUpdatedIAAScreen() {
	a.courses = []Course{}

	entryName := widget.NewEntry()
	entryName.SetPlaceHolder("Course name")
	entryCredits := widget.NewEntry()
	entryCredits.SetPlaceHolder("Credits")
	entryGrade := widget.NewEntry()
	entryGrade.SetPlaceHolder("Grade (0-10)")

	a.coursesTable = a.newCoursesTable(entryName, entryCredits, entryGrade)

	btnAdd := widget.NewButton("➕ Add", func() {
		a.addCourse(entryName, entryCredits, entryGrade)
	})

	hintLabel := widget.NewLabel("")
	hintLabel.Alignment = fyne.TextAlignCenter
	hintLabel.TextStyle = fyne.TextStyle{Italic: true}

	btnCalculate := widget.NewButton("✅ Calculate IAA", nil)
	btnClear := widget.NewButton("🗑️ Clear", func() {
		a.clearUpdatedFields(entryName, entryCredits, entryGrade)
		updateCalculateState(btnCalculate, hintLabel, a.currentIAA, a.completedCredits)
	})
	btnBack := widget.NewButton("⬅️ Back", a.createHomeScreen)

	a.resultLabel = widget.NewLabel("")
	a.resultLabel.Alignment = fyne.TextAlignCenter
	a.resultLabel.Wrapping = fyne.TextWrapWord

	update := func() {
		updateCalculateState(btnCalculate, hintLabel, a.currentIAA, a.completedCredits)
	}
	btnCalculate.OnTapped = a.calculateUpdatedIAA
	a.currentIAA.OnChanged = func(string) { update() }
	a.completedCredits.OnChanged = func(string) { update() }
	update()

	tableScroll := container.NewVScroll(a.coursesTable)
	tableScroll.SetMinSize(fyne.NewSize(0, 240))

	formFrame := container.NewVBox(
		a.createCard("Base Data", container.NewVBox(
			a.fieldWithLabel("Current IAA (0-10):", a.currentIAA),
			a.fieldWithLabel("Completed credits:", a.completedCredits),
		)),
		a.createCard("➕ Add Courses", container.NewVBox(
			container.NewBorder(nil, nil, nil, btnAdd,
				container.NewGridWithColumns(3, entryName, entryCredits, entryGrade),
			),
			widget.NewSeparator(),
			tableScroll,
		)),
	)

	title := widget.NewLabelWithStyle("📈 Calculate Updated IAA", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	subtitle := widget.NewLabel("Enter your info and add your semester courses.")
	subtitle.Alignment = fyne.TextAlignCenter
	subtitle.Wrapping = fyne.TextWrapWord

	content := container.NewVBox(
		title,
		subtitle,
		formFrame,
		container.NewGridWithColumns(3, btnCalculate, btnClear, btnBack),
		hintLabel,
		a.resultLabel,
	)

	a.window.SetContent(container.NewVScroll(content))
}

func (a *App) newCoursesTable(entryName, entryCredits, entryGrade *widget.Entry) *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return len(a.courses) + 1, 4 },
		newCourseCell,
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			a.updateCourseCell(id, obj.(*fyne.Container), entryName, entryCredits, entryGrade)
		},
	)
	table.SetColumnWidth(0, 360)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 140)
	return table
}

func newCourseCell() fyne.CanvasObject {
	label := widget.NewLabel("")
	label.Wrapping = fyne.TextWrapWord

	btnEdit := widget.NewButton("✏️", nil)
	btnRemove := widget.NewButton("🗑️", nil)
	actions := container.NewHBox(layout.NewSpacer(), btnEdit, btnRemove)

	return container.NewStack(label, actions)
}

func (a *App) updateCourseCell(id widget.TableCellID, cell *fyne.Container, entryName, entryCredits, entryGrade *widget.Entry) {
	label := cell.Objects[0].(*widget.Label)
	actions := cell.Objects[1].(*fyne.Container)
	btnEdit := actions.Objects[1].(*widget.Button)
	btnRemove := actions.Objects[2].(*widget.Button)

	if id.Row == 0 {
		label.Show()
		actions.Hide()
		label.TextStyle = fyne.TextStyle{Bold: true}
		label.Alignment = courseColumnAlignment(id.Col)
		label.SetText(courseHeader(id.Col))
		btnEdit.OnTapped = nil
		btnRemove.OnTapped = nil
		return
	}

	idx := id.Row - 1
	if idx < 0 || idx >= len(a.courses) {
		label.TextStyle = fyne.TextStyle{}
		label.Alignment = fyne.TextAlignLeading
		label.SetText("")
		label.Show()
		actions.Hide()
		btnEdit.OnTapped = nil
		btnRemove.OnTapped = nil
		return
	}

	if id.Col == 3 {
		label.SetText("")
		label.Hide()
		actions.Show()
		btnEdit.OnTapped = func() {
			a.editCourse(idx, entryName, entryCredits, entryGrade)
		}
		btnRemove.OnTapped = func() {
			a.removeCourse(idx)
		}
		return
	}

	label.TextStyle = fyne.TextStyle{}
	label.Alignment = courseColumnAlignment(id.Col)
	label.SetText(courseValue(a.courses[idx], id.Col))
	label.Show()
	actions.Hide()
	btnEdit.OnTapped = nil
	btnRemove.OnTapped = nil
}

func courseHeader(col int) string {
	switch col {
	case 0:
		return "Name"
	case 1:
		return "Credits"
	case 2:
		return "Grade"
	case 3:
		return "Actions"
	default:
		return ""
	}
}

func courseColumnAlignment(col int) fyne.TextAlign {
	if col == 0 {
		return fyne.TextAlignLeading
	}
	return fyne.TextAlignCenter
}

func courseValue(course Course, col int) string {
	switch col {
	case 0:
		return course.Name
	case 1:
		return strconv.Itoa(course.Credits)
	case 2:
		return fmt.Sprintf("%.2f", course.Grade)
	default:
		return ""
	}
}

func updateCalculateState(btn *widget.Button, hint *widget.Label, entries ...*widget.Entry) {
	if allFilled(entries...) {
		btn.Enable()
		if hint != nil {
			hint.SetText("")
		}
		return
	}

	btn.Disable()
	if hint != nil {
		hint.SetText("Fill in current IAA and credits to calculate.")
	}
}

func (a *App) addCourse(entryName, entryCredits, entryGrade *widget.Entry) {
	name := strings.TrimSpace(entryName.Text)
	creditsStr := strings.TrimSpace(entryCredits.Text)
	gradeStr := strings.TrimSpace(entryGrade.Text)

	if name == "" {
		dialog.ShowError(fmt.Errorf("please enter the course name"), a.window)
		return
	}

	credits, err := strconv.Atoi(creditsStr)
	if err != nil || credits <= 0 {
		dialog.ShowError(fmt.Errorf("credits must be a positive integer"), a.window)
		return
	}

	grade, err := strconv.ParseFloat(gradeStr, 64)
	if err != nil {
		dialog.ShowError(fmt.Errorf("invalid grade"), a.window)
		return
	}

	if grade < 0 || grade > 10 {
		dialog.ShowError(fmt.Errorf("grade must be between 0 and 10"), a.window)
		return
	}

	a.courses = append(a.courses, Course{Name: name, Credits: credits, Grade: grade})
	entryName.SetText("")
	entryCredits.SetText("")
	entryGrade.SetText("")
	if a.coursesTable != nil {
		a.coursesTable.Refresh()
	}
}

func (a *App) removeCourse(id int) {
	if id < 0 || id >= len(a.courses) {
		return
	}

	course := a.courses[id]
	dialog.ShowConfirm("Confirm Removal",
		fmt.Sprintf("Remove this course?\n\n%s - %d credits → Grade: %.2f", course.Name, course.Credits, course.Grade),
		func(confirmed bool) {
			if confirmed {
				a.courses = append(a.courses[:id], a.courses[id+1:]...)
				if a.coursesTable != nil {
					a.coursesTable.Refresh()
				}
			}
		}, a.window)
}

func (a *App) editCourse(id int, entryName, entryCredits, entryGrade *widget.Entry) {
	if id < 0 || id >= len(a.courses) {
		return
	}

	course := a.courses[id]
	entryName.SetText(course.Name)
	entryCredits.SetText(strconv.Itoa(course.Credits))
	entryGrade.SetText(fmt.Sprintf("%.2f", course.Grade))

	a.courses = append(a.courses[:id], a.courses[id+1:]...)
	if a.coursesTable != nil {
		a.coursesTable.Refresh()
	}
}

func (a *App) calculateUpdatedIAA() {
	currentIAA, completedCredits, err := a.validateBaseInputs()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	updatedIAA, err := CalculateUpdatedIAA(currentIAA, completedCredits, a.courses)
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	a.resultLabel.SetText(fmt.Sprintf("✨ Your updated IAA is: %.2f", updatedIAA))
}

func (a *App) clearUpdatedFields(entryName, entryCredits, entryGrade *widget.Entry) {
	a.currentIAA.SetText("")
	a.completedCredits.SetText("")
	a.courses = []Course{}
	if a.coursesTable != nil {
		a.coursesTable.Refresh()
	}
	a.resultLabel.SetText("")
	entryName.SetText("")
	entryCredits.SetText("")
	entryGrade.SetText("")
}
