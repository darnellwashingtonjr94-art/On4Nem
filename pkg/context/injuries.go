package context

type PlayerStatus struct {
	Name        string
	Designation string // "Probable", "Questionable", "Doubtful", "Out"
	WARValue    float64
}

func CalculateWAIL(players []PlayerStatus) float64 {
	totalLost := 0.0
	for _, p := range players {
		probability := 0.0
		switch p.Designation {
		case "Probable":
			probability = 0.05
		case "Questionable":
			probability = 0.27
		case "Doubtful":
			probability = 0.99
		case "Out":
			probability = 1.00
		}
		totalLost += p.WARValue * probability
	}
	return totalLost
}
