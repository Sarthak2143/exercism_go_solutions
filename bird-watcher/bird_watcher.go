package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, i := range birdsPerDay {
		sum += i
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	slice := birdsPerDay[((week-1)*7):((week-1)*7)+7]
	for _, i := range slice {
		sum += i
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i,_ := range birdsPerDay {
		if i%2==0{
			birdsPerDay[i] += 1
		} else {
			continue
		}
	}
	return birdsPerDay
}
