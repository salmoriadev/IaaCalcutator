package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// App representa a aplicação GUI
type App struct {
	app              fyne.App
	window           fyne.Window
	iaaAtual         *widget.Entry
	creditosCursados *widget.Entry
	materias         []Materia
	listaMaterias    *widget.List
	labelResultado   *widget.Label
}

// NovaApp cria uma nova instância da aplicação
func NovaApp() *App {
	myApp := app.NewWithID("iaa.calculator")
	window := myApp.NewWindow("📊 Calculadora de IAA")
	window.Resize(fyne.NewSize(850, 750))
	window.CenterOnScreen()

	return &App{
		app:      myApp,
		window:   window,
		materias: []Materia{},
	}
}

// Iniciar mostra a tela inicial e inicia o loop da aplicação
func (a *App) Iniciar() {
	a.criarTelaInicial()
	a.window.ShowAndRun()
}

// criarTelaInicial cria a tela inicial com menu
func (a *App) criarTelaInicial() {
	a.iaaAtual = widget.NewEntry()
	a.iaaAtual.SetPlaceHolder("Ex: 8.5")
	a.creditosCursados = widget.NewEntry()
	a.creditosCursados.SetPlaceHolder("Ex: 120")

	btnCalcularIAA := widget.NewButton("📈 Calcular IAA Atualizado", a.telaCalcularIAAAtualizado)
	btnCalcularMeta := widget.NewButton("🎯 Calcular Meta de IAA", a.telaCalcularMetaIAA)
	btnSair := widget.NewButton("🚪 Sair", a.window.Close)

	content := container.NewVBox(
		widget.NewLabelWithStyle("📊 Calculadora de IAA", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Escolha uma opção:"),
		widget.NewSeparator(),
		btnCalcularIAA,
		btnCalcularMeta,
		widget.NewSeparator(),
		btnSair,
	)

	a.window.SetContent(container.NewCenter(content))
}

// telaCalcularIAAAtualizado cria a tela para calcular IAA atualizado
func (a *App) telaCalcularIAAAtualizado() {
	a.materias = []Materia{}

	entryNome := widget.NewEntry()
	entryNome.SetPlaceHolder("Nome da Matéria")
	entryCreditos := widget.NewEntry()
	entryCreditos.SetPlaceHolder("Créditos")
	entryNota := widget.NewEntry()
	entryNota.SetPlaceHolder("Nota (0-10)")

	// Lista de matérias
	a.listaMaterias = widget.NewList(
		func() int { return len(a.materias) },
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel(""),
				widget.NewButton("✏️", nil),
				widget.NewButton("🗑️", nil),
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			box := obj.(*fyne.Container)
			label := box.Objects[0].(*widget.Label)
			m := a.materias[id]
			label.SetText(fmt.Sprintf("%d. %s - %d créditos → Nota: %.2f", id+1, m.Nome, m.Creditos, m.Nota))

			btnEditar := box.Objects[1].(*widget.Button)
			btnEditar.OnTapped = func() {
				a.editarMateria(id, entryNome, entryCreditos, entryNota)
			}

			btnRemover := box.Objects[2].(*widget.Button)
			btnRemover.OnTapped = func() {
				a.removerMateria(id)
			}
		},
	)

	btnAdicionar := widget.NewButton("➕ Adicionar", func() {
		a.adicionarMateria(entryNome, entryCreditos, entryNota)
	})
	btnCalcular := widget.NewButton("✅ Calcular IAA", a.calcularIAAAtualizado)
	btnLimpar := widget.NewButton("🗑️ Limpar", func() {
		a.limparCamposIAA(entryNome, entryCreditos, entryNota)
	})
	btnVoltar := widget.NewButton("⬅️ Voltar", a.criarTelaInicial)

	a.labelResultado = widget.NewLabel("")
	a.labelResultado.Alignment = fyne.TextAlignCenter

	formFrame := container.NewVBox(
		a.criarCard("Dados Iniciais", container.NewVBox(
			a.criarCampoComLabel("IAA Atual (0-10):", a.iaaAtual),
			a.criarCampoComLabel("Créditos já cursados:", a.creditosCursados),
		)),
		a.criarCard("➕ Adicionar Matérias", container.NewVBox(
			container.NewHBox(entryNome, entryCreditos, entryNota, btnAdicionar),
			widget.NewSeparator(),
			container.NewScroll(a.listaMaterias),
		)),
	)

	content := container.NewVBox(
		widget.NewLabelWithStyle("📈 Calcular IAA Atualizado", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		formFrame,
		container.NewHBox(btnCalcular, btnLimpar, btnVoltar),
		a.labelResultado,
	)

	a.window.SetContent(container.NewScroll(content))
}

// telaCalcularMetaIAA cria a tela para calcular meta de IAA
func (a *App) telaCalcularMetaIAA() {
	entryCreditosSemestre := widget.NewEntry()
	entryCreditosSemestre.SetPlaceHolder("Créditos do semestre atual")
	entryIAADesejado := widget.NewEntry()
	entryIAADesejado.SetPlaceHolder("IAA Desejado (0-10)")

	labelResultadoMeta := widget.NewLabel("")
	labelResultadoMeta.Alignment = fyne.TextAlignCenter

	btnCalcular := widget.NewButton("🎯 Calcular Meta", func() {
		a.calcularMetaIAA(entryCreditosSemestre, entryIAADesejado, labelResultadoMeta)
	})
	btnLimpar := widget.NewButton("🗑️ Limpar", func() {
		a.limparCamposMeta(entryCreditosSemestre, entryIAADesejado, labelResultadoMeta)
	})
	btnVoltar := widget.NewButton("⬅️ Voltar", a.criarTelaInicial)

	formFrame := a.criarCard("Dados", container.NewVBox(
		a.criarCampoComLabel("IAA Atual (0-10):", a.iaaAtual),
		a.criarCampoComLabel("Créditos já cursados:", a.creditosCursados),
		a.criarCampoComLabel("Créditos do semestre atual:", entryCreditosSemestre),
		a.criarCampoComLabel("IAA Desejado (0-10):", entryIAADesejado),
	))

	content := container.NewVBox(
		widget.NewLabelWithStyle("🎯 Calcular Meta de IAA", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		formFrame,
		container.NewHBox(btnCalcular, btnLimpar, btnVoltar),
		labelResultadoMeta,
	)

	a.window.SetContent(container.NewScroll(content))
}

// Métodos auxiliares para criar componentes da UI
func (a *App) criarCard(titulo string, conteudo fyne.CanvasObject) *widget.Card {
	return widget.NewCard(titulo, "", conteudo)
}

func (a *App) criarCampoComLabel(labelText string, entry *widget.Entry) fyne.CanvasObject {
	return container.NewBorder(nil, nil, widget.NewLabel(labelText), nil, entry)
}

// Métodos para manipulação de matérias
func (a *App) adicionarMateria(entryNome, entryCreditos, entryNota *widget.Entry) {
	nome := entryNome.Text
	creditosStr := entryCreditos.Text
	notaStr := entryNota.Text

	if nome == "" {
		dialog.ShowError(fmt.Errorf("por favor, informe o nome da matéria"), a.window)
		return
	}

	creditos, err := strconv.Atoi(creditosStr)
	if err != nil || creditos <= 0 {
		dialog.ShowError(fmt.Errorf("créditos devem ser um número inteiro positivo"), a.window)
		return
	}

	nota, err := strconv.ParseFloat(notaStr, 64)
	if err != nil {
		dialog.ShowError(fmt.Errorf("nota inválida"), a.window)
		return
	}

	if nota < 0 || nota > 10 {
		dialog.ShowError(fmt.Errorf("nota deve estar entre 0 e 10"), a.window)
		return
	}

	a.materias = append(a.materias, Materia{Nome: nome, Creditos: creditos, Nota: nota})
	entryNome.SetText("")
	entryCreditos.SetText("")
	entryNota.SetText("")
	a.listaMaterias.Refresh()
}

func (a *App) removerMateria(id int) {
	if id < 0 || id >= len(a.materias) {
		return
	}

	m := a.materias[id]
	dialog.ShowConfirm("Confirmar Remoção",
		fmt.Sprintf("Deseja remover a matéria?\n\n%s - %d créditos → Nota: %.2f", m.Nome, m.Creditos, m.Nota),
		func(confirmado bool) {
			if confirmado {
				a.materias = append(a.materias[:id], a.materias[id+1:]...)
				a.listaMaterias.Refresh()
			}
		}, a.window)
}

func (a *App) editarMateria(id int, entryNome, entryCreditos, entryNota *widget.Entry) {
	if id < 0 || id >= len(a.materias) {
		return
	}

	m := a.materias[id]
	entryNome.SetText(m.Nome)
	entryCreditos.SetText(strconv.Itoa(m.Creditos))
	entryNota.SetText(fmt.Sprintf("%.2f", m.Nota))

	// Remove a matéria antiga para re-adicionar editada
	a.materias = append(a.materias[:id], a.materias[id+1:]...)
	a.listaMaterias.Refresh()
}

// Métodos para cálculos
func (a *App) calcularIAAAtualizado() {
	iaa, creditos, err := a.validarDadosIniciais()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	iaaAtualizado, err := CalcularIAAAtualizado(iaa, creditos, a.materias)
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	a.labelResultado.SetText(fmt.Sprintf("✨ Seu IAA atualizado é: %.2f", iaaAtualizado))
}

func (a *App) calcularMetaIAA(entryCreditosSemestre, entryIAADesejado *widget.Entry, labelResultado *widget.Label) {
	iaa, creditosCursados, err := a.validarDadosIniciais()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	creditosSemestreStr := entryCreditosSemestre.Text
	creditosSemestre, err := strconv.Atoi(creditosSemestreStr)
	if err != nil || creditosSemestre <= 0 {
		dialog.ShowError(fmt.Errorf("créditos do semestre devem ser um número inteiro positivo"), a.window)
		return
	}

	objetivoStr := entryIAADesejado.Text
	objetivo, err := strconv.ParseFloat(objetivoStr, 64)
	if err != nil {
		dialog.ShowError(fmt.Errorf("IAA desejado inválido"), a.window)
		return
	}

	if objetivo < 0 || objetivo > 10 {
		dialog.ShowError(fmt.Errorf("IAA desejado deve estar entre 0 e 10"), a.window)
		return
	}

	mediaNecessaria, pontosFaltam, err := CalcularMetaIAA(iaa, creditosCursados, creditosSemestre, objetivo)
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	var resultado string
	if mediaNecessaria > 10 {
		resultado = "❌ Não é possível alcançar esse IAA com os créditos planejados."
	} else if mediaNecessaria < 0 {
		resultado = fmt.Sprintf("✅ Você já atingiu o IAA %.2f!\n(média necessária: %.2f)", objetivo, mediaNecessaria)
	} else {
		resultado = fmt.Sprintf("🎯 Para alcançar IAA %.2f:\n\n📊 Média necessária: %.2f\n📈 Acumulado necessário: %.2f",
			objetivo, mediaNecessaria, pontosFaltam)
	}

	labelResultado.SetText(resultado)
}

// Métodos auxiliares
func (a *App) validarDadosIniciais() (float64, int, error) {
	iaaStr := a.iaaAtual.Text
	iaa, err := strconv.ParseFloat(iaaStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("IAA inválido")
	}

	if iaa < 0 || iaa > 10 {
		return 0, 0, fmt.Errorf("IAA deve estar entre 0 e 10")
	}

	creditosStr := a.creditosCursados.Text
	creditos, err := strconv.Atoi(creditosStr)
	if err != nil {
		return 0, 0, fmt.Errorf("créditos inválidos")
	}

	if creditos < 0 {
		return 0, 0, fmt.Errorf("créditos devem ser não negativos")
	}

	return iaa, creditos, nil
}

func (a *App) limparCamposIAA(entryNome, entryCreditos, entryNota *widget.Entry) {
	a.iaaAtual.SetText("")
	a.creditosCursados.SetText("")
	a.materias = []Materia{}
	if a.listaMaterias != nil {
		a.listaMaterias.Refresh()
	}
	a.labelResultado.SetText("")
	entryNome.SetText("")
	entryCreditos.SetText("")
	entryNota.SetText("")
}

func (a *App) limparCamposMeta(entryCreditosSemestre, entryIAADesejado *widget.Entry, labelResultado *widget.Label) {
	a.iaaAtual.SetText("")
	a.creditosCursados.SetText("")
	entryCreditosSemestre.SetText("")
	entryIAADesejado.SetText("")
	labelResultado.SetText("")
}
