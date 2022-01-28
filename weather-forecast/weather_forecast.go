// Package weather is a tool that provides
// information about forecast.
package weather

// CurrentCondition represnets information about condition.
var CurrentCondition string

// CurrentLocation represents information about location.
var CurrentLocation string

// Forecast returns string that contains current location and weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
