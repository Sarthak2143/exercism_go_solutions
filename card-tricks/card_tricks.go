package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	list := []int{2, 6, 9}
	return list
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index > (len(slice)-1) || index < 0 {
		return -1
	}
	item := slice[index]
	return item
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index > (len(slice)-1) || index < 0 {
		new_slice := append(slice, value)
		return new_slice
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	if len(value) == 0 {
		return slice
	}
	new_slice := append(value, slice...)
	return new_slice
}
func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}


// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < (len(slice)){
		new_slice := remove(slice, index)
		return new_slice
	}
	return slice
}
