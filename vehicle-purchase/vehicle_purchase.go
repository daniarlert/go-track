package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	bestOption := option2
	if option1 < option2 {
		bestOption = option1
	}

	return fmt.Sprintf("%s is clearly the better choice.", bestOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// less than 3 = 80% original price
	// at least 10 = 50%
	// least 3 but less than 10 = 70%
	if age < 3 {
		return originalPrice - (originalPrice * 0.20)
	} else if age >= 3 && age < 10 {
		return originalPrice - (originalPrice * 0.30)
	} else {
		return originalPrice / 2
	}
}
