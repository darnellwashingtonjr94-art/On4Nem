package portfolio

import "log"

type ExposureTracker struct {
	Limits map[string]float64
}

func NewExposureTracker() *ExposureTracker {
	return &ExposureTracker{Limits: make(map[string]float64)}
}

func (et *ExposureTracker) TrackLeagueExposure(league string, stakeAmount float64) {
	et.Limits[league] += stakeAmount
	log.Printf("Exposure updated for league [%s]: Total Capital at Risk = $%.2f", league, et.Limits[league])
}
