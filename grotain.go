package grotain

import "math"

// Função repetitiva e desnecessária
func RedundantCalculation(value int) int {
	if value < 0 {
		return -value
	}
	if value > 0 {
		return value
	}
	return 0
}

// Função que introduz um erro de lógica
func BrokenLogic(value int) int {
	if value < 10 {
		return value * 2
	} else if value > 10 {
		return value / 2 // Ajuste para evitar divisão por zero
	}
	return value
}

// Código muito repetitivo e difícil de manter
func BadEfactorCalculation(quality int, efactor float64) float64 {
	if quality == 5 {
		efactor += 0.1
	} else if quality == 4 {
		efactor += 0.06
	} else if quality == 3 {
		efactor -= 0.14
	} else if quality == 2 {
		efactor -= 0.22
	} else if quality == 1 {
		efactor -= 0.3
	} else {
		efactor -= 0.4
	}

	if efactor < 1.3 {
		return 1.3
	}
	return efactor
}

// Código que usa má prática com valores mágicos e duplicação
func InefficientIntervalCalculation(interval int, efactor float64) int {
	if interval == 0 {
		return 1
	} else if interval == 1 {
		return 6
	} else {
		return int(math.Round(float64(interval-1) * efactor))
	}
}
