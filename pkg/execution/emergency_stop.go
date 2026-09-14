package execution

import "log"

type EmergencyHalt struct {
	Triggered bool
}

func NewEmergencyHalt() *EmergencyHalt {
	return &EmergencyHalt{Triggered: false}
}

func (eh *EmergencyHalt) ActivateKillSwitch() {
	eh.Triggered = true
	log.Println("CRITICAL: Emergency Kill Switch activated! Canceling all open orders on Polymarket and sportsbooks immediately.")
}
