package main

import (
	"flag"
	"log"
)

func HandleCLIArguments() {
	mode := flag.String("mode", "auto", "Operating mode: auto, manual, backtest")
	flag.Parse()
	log.Printf("On4Nem CLI initialized in mode: %s", *mode)
}
