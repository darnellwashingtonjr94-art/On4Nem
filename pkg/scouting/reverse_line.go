package scouting

type LineMovement struct {
	GameID       string
	PublicPct    float64 // Public betting percentage on favorite
	OpeningSpread float64
	CurrentSpread float64
}

func DetectReverseLineMovement(lm LineMovement) bool {
	// If public is heavy on favorite (>75%), but spread moves toward underdog, it's RLM (sharp money)
	if lm.PublicPct > 0.75 && lm.CurrentSpread < lm.OpeningSpread {
		return true
	}
	return false
}
