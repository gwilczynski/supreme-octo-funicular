package lasagna

const averagePreparationTime int = 2

// PreparationTime returns the estimate for the total preparation time
// based on the number of layers
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * averagePreparationTime
	}

	return len(layers) * time
}

const noodlesQuantity int = 50
const sauceQuantity float64 = 0.2

// Quantities determine the quantity of noodles and sauce needed to make your meal
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0

	for _, item := range layers {
		switch item {
		case "noodles":
			noodles += noodlesQuantity
		case "sauce":
			sauce += sauceQuantity
		}
	}

	return
}

func lastElementIndex(list []string) int {
	return len(list) - 1
}

// AddSecretIngredient
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[lastElementIndex(myList)] = friendsList[lastElementIndex(friendsList)]
}

// ScaleRecipe
func ScaleRecipe(inputList []float64, portions int) (slice []float64) {
	slice = make([]float64, len(inputList))

	for index, quantity := range inputList {
		slice[index] = quantity / 2 * float64(portions)
	}

	return

}
