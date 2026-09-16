package main

import (
	"fmt"
	"log"

	"github.com/darnellwashingtonjr94-art/On4Nem/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Printf("Starting On4Nem Engine in %s mode\n", cfg.Environment)
	log.Println("Core agent initialized successfully.")
}
