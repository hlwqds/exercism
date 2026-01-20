package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if v, exists := units[unit]; !exists {
		return false
	} else {
		bill[item] += v
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	c, exists := bill[item]
	if !exists {
		return false
	}
	v, exists := units[unit]
	if !exists {
		return false
	}
	if c-v < 0 {
		return false
	}
	if c-v == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = c - v
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, exists := bill[item]
	if !exists {
		return 0, false
	}

	return v, true
}
