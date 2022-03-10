package cars

const CostPerTenCars = 95_000
const CostPerIndividualCar = 10_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate / 100) * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(carsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	numberOfCars := carsCount / 10
	carsRemaining := carsCount % 10

	return uint(numberOfCars)*CostPerTenCars + uint(carsRemaining)*CostPerIndividualCar
}
