package scouting

import "log"

type SteamAlert struct {
	GameID    string
	BookCount int
	Direction string
}

func AggregateSteamMoves(marketShifts map[string]float64) []SteamAlert {
	var alerts []SteamAlert
	for gameID, shift := range marketShifts {
		if shift >= 0.5 { // Sharp line movement across multiple books simultaneously
			log.Printf("STEAM MOVE DETECTED on Game ID: %s", gameID)
			alerts = append(alerts, SteamAlert{GameID: gameID, BookCount: 5, Direction: "Sharp Money"})
		}
	}
	return alerts
}
