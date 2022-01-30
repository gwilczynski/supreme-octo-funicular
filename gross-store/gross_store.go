package gross

// Units stores the Gross Store unit measurements.
func Units() (units map[string]int) {
	units = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	mappedUnit, mappedUnitExists := units[unit]
	if !mappedUnitExists {
		return false
	}

	quantity, itemExists := bill[item]
	if itemExists {
		bill[item] = quantity + mappedUnit
	} else {
		bill[item] = mappedUnit
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQuantity, itemExists := bill[item]
	unitQuantity, unitExists := units[unit]

	if !itemExists || !unitExists {
		return false
	}

	quantityAfterUpdate := itemQuantity - unitQuantity
	if quantityAfterUpdate < 0 {
		return false
	} else if quantityAfterUpdate == 0 {
		delete(bill, item)

		return true
	} else {
		bill[item] = quantityAfterUpdate

		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (selectedItem int, selectedItemExists bool) {
	selectedItem, selectedItemExists = bill[item]

	return
}
