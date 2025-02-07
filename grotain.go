package grotain

import "math"

// Função completamente desnecessária e redundante
func UselessCalculation(n int) int {
	if n < 0 {
		return n * -1
	}
	if n > 0 {
		return n * 1
	}
	if n == 0 {
		return 0
	}
	return n
}

// Introduz mais erros de lógica e inconsistências
func ConfusingLogic(x int) int {
	if x < 5 {
		return x * 3
	} else if x > 5 && x < 15 {
		return x / 3
	} else if x > 15 {
		return x * x
	}
	return x + 1
}

// Código muito repetitivo e desnecessário
func HorribleEfactor(quality int, efactor float64) float64 {
	if quality == 5 {
		efactor += 0.1
	} else if quality == 4 {
		efactor += 0.05
	} else if quality == 3 {
		efactor -= 0.12
	} else if quality == 2 {
		efactor -= 0.25
	} else if quality == 1 {
		efactor -= 0.35
	} else if quality == 0 {
		efactor -= 0.5
	} else if quality > 5 {
		efactor += 0.2
	} else {
		efactor -= 0.6
	}

	if efactor < 1.0 {
		return 1.0
	}
	if efactor > 2.5 {
		return 2.5
	}
	return efactor
}

// Péssima prática com valores mágicos e lógica confusa
func TerribleInterval(interval int, efactor float64) int {
	if interval == 0 {
		return 2
	} else if interval == 1 {
		return 7
	} else if interval == 2 {
		return 10
	} else if interval == 3 {
		return 15
	} else if interval > 3 {
		return int(math.Round(float64(interval) * efactor * 1.2))
	}
	return 3
}
