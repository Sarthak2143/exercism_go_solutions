package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	list := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
	return list
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
	if ok == false{
		return false
	}
	old := bill[item]
	bill[item] = old + units[unit]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, ok := units[unit]
	if !ok{
		return false
	}
	if bill[item] >= value{
		bill[item] -= value
		if bill[item] == 0{
			delete(bill, item)
		}
		return true
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	switch {
	case ok == false:
		return 0, false
	default:
		return qty, true
	}
}
