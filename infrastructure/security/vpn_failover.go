package security

import (
	"log"
	"net/http"
	"os/exec"
	"time"
)

type VPNMonitor struct {
	InterfaceName string
	CheckInterval time.Duration
}

func NewVPNMonitor(iface string, interval time.Duration) *VPNMonitor {
	return &VPNMonitor{InterfaceName: iface, CheckInterval: interval}
}

func (vm *VPNMonitor) StartVPNWatchdog() {
	ticker := time.NewTicker(vm.CheckInterval)
	go func() {
		for range ticker.C {
			client := &http.Client{Timeout: 3 * time.Second}
			resp, err := client.Get("https://api.ipify.org")
			if err != nil {
				log.Println("[VPN WARNING] Public IP check failed. VPN tunnel may be down. Reconnecting...")
				vm.RestartVPN()
			} else {
				_ = resp.Body.Close()
			}
		}
	}()
}

func (vm *VPNMonitor) RestartVPN() {
	log.Printf("[VPN] Restarting secure interface %s...", vm.InterfaceName)
	cmd := exec.Command("wg-quick", "down", vm.InterfaceName)
	_ = cmd.Run()
	time.Sleep(1 * time.Second)
	upCmd := exec.Command("wg-quick", "up", vm.InterfaceName)
	_ = upCmd.Run()
}
