package main

import (
	"log"
	"On4Nem/config"
)

type Bootstrapper struct {
	Config *config.Config
}

func NewBootstrapper(cfg *config.Config) *Bootstrapper {
	return &Bootstrapper{Config: cfg}
}

func (b *Bootstrapper) InitializeSubsystems() {
	log.Println("Bootstrapping On4Nem infrastructure dependencies, storage buckets, and RPC bindings...")
}
