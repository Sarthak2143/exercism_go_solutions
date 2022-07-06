package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	cars := float64(productionRate)
	percentage := successRate/100
	amount := cars*percentage
	return amount

}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	inHours := CalculateWorkingCarsPerHour(productionRate, successRate)
	inMins := int(inHours/60)
	return inMins
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	cars := carsCount
	cars_1 := (cars/10)*10
	cars_2 := cars%10
	cost := uint(cars_1*9500 + cars_2*10000)
	return cost
}
