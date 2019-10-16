package space

type Planet string

const earthPeriod = 31557600

var planetPeriod = map[Planet]float64{
	"Earth":   earthPeriod,
	"Mercury": earthPeriod * 0.2408467,
	"Venus":   earthPeriod * 0.61519726,
	"Mars":    earthPeriod * 1.8808158,
	"Jupiter": earthPeriod * 11.862615,
	"Saturn":  earthPeriod * 29.447498,
	"Uranus":  earthPeriod * 84.016846,
	"Neptune": earthPeriod * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / planetPeriod[planet]
}
