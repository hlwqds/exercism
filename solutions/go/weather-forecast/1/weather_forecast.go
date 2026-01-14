// Package weather comment.
package weather

var (
	// CurrentCondition is comment.
	CurrentCondition string
	// CurrentLocation is comment.
	CurrentLocation string
)

// Forecast is comment.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
