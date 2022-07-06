package speed

// TODO: define the 'Car' type struct

type Car struct{
	speed, batteryDrain, battery, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed, batteryDrain, 100, 0}
}

// TODO: define the 'Track' type struct

type Track struct{
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery > car.batteryDrain{
		return Car{car.speed, car.batteryDrain, car.battery-(car.batteryDrain), car.speed+car.distance}
	}
	return Car{car.speed, car.batteryDrain, car.battery, car.distance}

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	sec := track.distance/car.speed
	drain := car.batteryDrain * sec
	if car.battery - drain >= 0{
		return true
	}
	return false
}
