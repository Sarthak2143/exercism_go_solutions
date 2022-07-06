package lasagna

const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	remTime := OvenTime - actualMinutesInOven
	return remTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	prepTime := numberOfLayers*2
	return prepTime
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	remTime := actualMinutesInOven
	prepTime := PreparationTime(numberOfLayers)
	time := remTime + prepTime
	return time
}
