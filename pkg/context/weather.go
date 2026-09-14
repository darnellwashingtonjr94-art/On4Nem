package context

type WeatherCondition struct {
	Temperature float64 // Fahrenheit
	WindSpeed   float64 // MPH
	Precipitation bool
	IsDome      bool
}

func AdjustTotalForWeather(baseTotal float64, w WeatherCondition) float64 {
	if w.IsDome {
		return baseTotal
	}
	adjusted := baseTotal
	if w.WindSpeed > 15.0 {
		adjusted -= (w.WindSpeed * 0.15) // Heavy wind lowers scoring totals (NFL/Soccer)
	}
	if w.Precipitation {
		adjusted *= 0.93 // 7% reduction for wet conditions
	}
	return adjusted
}
