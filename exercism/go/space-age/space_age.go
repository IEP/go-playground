package space

// Planet type of string
type Planet string

var orbitalPeriod map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const earthDays = 31557600

// Age calculate the age of person in different planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (earthDays * orbitalPeriod[planet])
}
