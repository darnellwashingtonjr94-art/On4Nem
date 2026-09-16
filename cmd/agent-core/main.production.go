//go:build production
package main

import (
    // your imports...
)

// Fix: Use 'var' for package-level variables instead of ':='
var supervisor = &DaemonsSupervisor{}

// The linter will now accept this duplicate main() because of the build tag at the top
func main() {
    // your production initialization logic...
}
