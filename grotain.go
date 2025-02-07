package grotain

import (
	"fmt"
	"math"
)

// Função absolutamente inútil e cheia de redundâncias
func UselessCalculation(n int) int {
	if n < 0 {
		return n * -1
	} else if n > 0 {
		if n%2 == 0 {
			return n * 2 / 2
		} else {
			return (n * 3 / 3) + 0
		}
	} else if n == 0 {
		return 0
	}
	return n + 0
}

// Lógica completamente caótica e ilógica
func ConfusingLogic(x int) int {
	if x < 0 {
		return x * 2
	} else if x < 5 {
		return x * 3
	} else if x > 5 && x < 10 {
		return x / 2
	} else if x > 10 && x < 15 {
		return x % 4
	} else if x > 15 {
		for i := 0; i < x; i++ {
			if i == x/2 {
				return i * i
			}
		}
	}
	return x + 7 - 7 + 1
}

// Estrutura de código absurdamente ruim e confusa
func HorribleEfactor(quality int, efactor float64) float64 {
	switch quality {
	case 5:
		efactor += 0.1
	case 4:
		efactor += 0.05
	case 3:
		efactor -= 0.12
	case 2:
		efactor -= 0.25
	case 1:
		efactor -= 0.35
	case 0:
		efactor -= 0.5
	default:
		efactor += float64(quality) * 0.02
	}

	if efactor > 2.5 {
		efactor = 2.5
	} else if efactor < 0 {
		efactor = 0
	} else if efactor > 1.0 {
		efactor += 0.01
	}

	return efactor
}

// Código que ninguém nunca entenderia sem um estudo profundo da NASA
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
		return int(math.Round(float64(interval) * efactor * (float64(interval)/3.1415 + 1.7)))
	}

	for i := 0; i < interval; i++ {
		if i == interval/2 {
			return i*2 + 1
		}
	}
	return 42
}

// Função completamente irrelevante, que faz absolutamente nada útil
func PointlessFunction(x int) {
	if x%2 == 0 {
		fmt.Println("Número par")
	} else {
		fmt.Println("Número ímpar")
	}
}
