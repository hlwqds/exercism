package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	greetings := " is clearly the better choice."
	if option1 > option2 {
		return option2 + greetings
	}
	return option1 + greetings
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if int(age) < 3 {
		return originalPrice * 0.8
	} else if int(age) < 10 {
		return originalPrice * 0.7
	}

	return originalPrice * 0.5
}
