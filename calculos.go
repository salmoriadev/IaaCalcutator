package main

import "fmt"

// CalcularIAAAtualizado calcula o IAA após adicionar novas matérias
func CalcularIAAAtualizado(iaaAtual float64, creditosCursados int, materias []Materia) (float64, error) {
	if creditosCursados < 0 {
		return 0, fmt.Errorf("créditos cursados devem ser não negativos")
	}

	if creditosCursados == 0 && len(materias) == 0 {
		return 0, fmt.Errorf("adicione créditos para calcular")
	}

	// Calcular pontos totais
	pontosTotais := iaaAtual * float64(creditosCursados)
	creditosTotais := creditosCursados

	for _, m := range materias {
		pontosTotais += m.Nota * float64(m.Creditos)
		creditosTotais += m.Creditos
	}

	if creditosTotais == 0 {
		return 0, fmt.Errorf("não há créditos registrados")
	}

	return pontosTotais / float64(creditosTotais), nil
}

// CalcularMetaIAA calcula a média necessária para alcançar um IAA objetivo
func CalcularMetaIAA(iaaAtual float64, creditosCursados int, creditosSemestre int, objetivo float64) (float64, float64, error) {
	if creditosSemestre <= 0 {
		return 0, 0, fmt.Errorf("créditos do semestre devem ser positivos")
	}

	totalCreditos := creditosCursados + creditosSemestre
	pontosAtuais := iaaAtual * float64(creditosCursados)
	pontosNecessarios := objetivo * float64(totalCreditos)
	pontosFaltam := pontosNecessarios - pontosAtuais
	mediaNecessaria := pontosFaltam / float64(creditosSemestre)

	return mediaNecessaria, pontosFaltam, nil
}
