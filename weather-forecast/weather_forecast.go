// Package weather provides tools to get the current
// forecast on different cities.
package weather

// CurrentCondition refers to the current condition of a city.
var CurrentCondition string

// CurrentLocation refers to the current location of the user.
var CurrentLocation string

// Forecast returns an string with the current location and a report
// about the weather condition. It recieves the city and a condition
// as parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
