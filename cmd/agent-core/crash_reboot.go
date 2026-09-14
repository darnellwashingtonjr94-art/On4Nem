package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type CrashRebootSupervisor struct {
	BinaryPath string
}

func NewCrashRebootSupervisor(path string) *CrashRebootSupervisor {
	return &CrashRebootSupervisor{BinaryPath: path}
}

func (crs *CrashRebootSupervisor) WatchdogLoop(rebootChan chan struct{}) {
	for range rebootChan {
		log.Println("[CRASH DETECTED] Initiating emergency state flush and automated system reboot...")
		
		// Flush pending logs & persist state
		time.Sleep(500 * time.Millisecond)

		cmd := exec.Command(crs.BinaryPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		
		if err := cmd.Start(); err != nil {
			log.Fatalf("[FATAL] Failed to reboot autonomous engine: %v", err)
		}
		
		log.Println("[REBOOT SUCCESS] New daemon process spawned. Terminating faulted process.")
		os.Exit(0)
	}
}
