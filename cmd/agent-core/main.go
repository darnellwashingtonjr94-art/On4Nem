package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"On4Nem/config"
	"On4Nem/pkg/agents"
	"On4Nem/pkg/execution"
	"On4Nem/pkg/treasury"
)

func main() {
	log.Println("Initializing On4Nem Autonomous Engine...")
	cfg := config.LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize Execution Clients & Monad RPC
	monadClient := execution.NewMonadRPC(cfg.MonadRPCURL)
	streamer := execution.NewStreamer(cfg.WSUrls)
	go streamer.Start(ctx)

	// Initialize Treasury & Circuit Breakers
	breaker := treasury.NewCircuitBreaker(cfg.MaxExposureLimit)

	// Initialize AI Manager Agent (Gemini Pro Thinking Core)
	manager := agents.NewManagerAgent(cfg.GeminiAPIKey, monadClient, breaker)
	go manager.RunOrchestrationLoop(ctx)

	// Graceful Shutdown Handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down On4Nem gracefully...")
	cancel()
}
