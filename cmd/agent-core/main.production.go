//go:build production

package main

import (
	"fmt"
	"log"

	"github.com/darnellwashingtonjr94-art/On4Nem/config"
)

type DaemonsSupervisor struct {
	Active bool
}

// Fix: Replaced invalid top-level ':=' with 'var'
var supervisor = &DaemonsSupervisor{Active: true}

// Fix: The //go:build production directive prevents duplicate main error during standard builds
func main() {
	cfg := config.LoadConfig()
	fmt.Printf("Starting On4Nem Production Engine [Supervisor Active: %v]\n", supervisor.Active)
	log.Printf("Production server listening on port %s\n", cfg.Port)
}
