// Package space package implements calculating age in various planet years.
package space

// Planet is a string type to represent planets.
type Planet string

// Age function calculates age of sameone in vaious planet years,
// it takes age in seconds and planet and returns age in that planet
// years.
func Age(seconds float64, planet Planet) float64 {
	const earthAge = 31557600
	var result float64

	switch planet {
	case "Earth":
		result = seconds / earthAge
	case "Mercury":
		result = seconds / (earthAge * 0.2408467)
	case "Venus":
		result = seconds / (earthAge * 0.61519726)
	case "Mars":
		result = seconds / (earthAge * 1.8808158)
	case "Jupiter":
		result = seconds / (earthAge * 11.862615)
	case "Saturn":
		result = seconds / (earthAge * 29.447498)
	case "Uranus":
		result = seconds / (earthAge * 84.016846)
	case "Neptune":
		result = seconds / (earthAge * 164.79132)
	}

	return result
}
