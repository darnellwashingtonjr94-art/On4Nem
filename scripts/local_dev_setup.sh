#!/bin/bash
# On4Nem Local Development Environment Bootstrapper
echo "Starting On4Nem local infrastructure (PostgreSQL, Prometheus, Redis)..."
docker-compose up -d postgres redis prometheus
echo "Infrastructure running. Initializing Go agent daemon..."
go run cmd/agent-core/main.go cmd/agent-core/bootstrap.go cmd/agent-core/healthcheck.go cmd/agent-core/aggregator.go cmd/agent-core/cli.go cmd/agent-core/shutdown.go cmd/agent-core/initializer.go cmd/agent-core/reporter.go cmd/agent-core/orchestrator_daemon.go
