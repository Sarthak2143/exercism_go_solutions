package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(slice []string, time int) int {
	if time == 0{
		return len(slice)*2
	}
	return len(slice)*time
}

// TODO: define the 'Quantities()' function

func Quantities(slice []string) (int, float64) {
	noodle := 0
	sauce := 0.0
	for _, i := range slice {
		if i == "noodles" {
			noodle += 1
		} else if i == "sauce"{
			sauce += 1
		} else {
			continue
		}
	}
	return noodle*50, sauce*0.20
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, ownList []string) {
	ownList[len(ownList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amounts int) []float64 {
    scale := float64(amounts) / 2.0
    var res []float64
    for i := 0; i < len(quantities); i++ {
        res = append(res, quantities[i] * float64(scale))
    }
    return res
}
