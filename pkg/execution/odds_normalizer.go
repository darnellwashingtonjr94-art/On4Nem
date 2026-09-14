package execution

import "fmt"

func NormalizeToDecimal(oddsFormat string, rawValue float64) (float64, error) {
	switch oddsFormat {
	case "decimal":
		return rawValue, nil
	case "american":
		if rawValue > 0 {
			return (rawValue / 100.0) + 1.0, nil
		}
		return (100.0 / -rawValue) + 1.0, nil
	case "fractional":
		return rawValue + 1.0, nil
	default:
		return 0, fmt.Errorf("unsupported odds format: %s", oddsFormat)
	}
}
