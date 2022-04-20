package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer <= 0 {
		return len(layers) * 2
	}

	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	// 50 grams of noodles for each each noodle layer
	// 0.2 liters of sauce for each sauce layer
	var noodles int
	var sauce float64

	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
			continue
		}

		if layer == "noodles" {
			noodles += 50
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(secretIngredients, ingredients []string) {
	ingredients[len(ingredients)-1] = secretIngredients[len(secretIngredients)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := make([]float64, len(amounts))

	for i, amount := range amounts {
		newAmounts[i] = amount * float64(portions) / 2
	}

	return newAmounts
}
