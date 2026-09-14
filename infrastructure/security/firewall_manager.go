package security

import (
	"log"
	"os/exec"
)

type FirewallManager struct {
	AllowedPorts []string
}

func NewFirewallManager(ports []string) *FirewallManager {
	return &FirewallManager{AllowedPorts: ports}
}

func (fm *FirewallManager) ApplyUFWHardening() error {
	log.Println("[Security] Applying automated UFW firewall hardening rules...")
	
	commands := []string{
		"ufw default deny incoming",
		"ufw default allow outgoing",
		"ufw allow 22/tcp", // SSH access
		"ufw allow 9090/tcp", // Prometheus metrics
	}

	for _, cmdStr := range commands {
		cmd := exec.Command("sh", "-c", cmdStr)
		if err := cmd.Run(); err != nil {
			log.Printf("[Warning] UFW command failed (%s): %v", cmdStr, err)
		}
	}

	enableCmd := exec.Command("ufw", "--force", "enable")
	return enableCmd.Run()
}
