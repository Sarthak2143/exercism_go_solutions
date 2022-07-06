package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	text := " is clearly the better choice."
	if option1 > option2 {
		return option2 + text
	} else if option2 > option1 {
		return option1 + text
	} else {
		return "ur mom gae"
	}

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		price := originalPrice*0.80
		return price
	} else if age >= 10 {
		price := originalPrice*0.50
		return price
	} else if age >= 3 && age < 10 {
		price := originalPrice*0.70
		return price
	} else {
		return 0
	}
}
