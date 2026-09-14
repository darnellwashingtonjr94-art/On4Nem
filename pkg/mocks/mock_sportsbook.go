package mocks

import (
	"log"
	"math/rand"
	"time"
)

type MockSportsbookServer struct {
	BookName string
	Port     string
}

func NewMockSportsbook(name string, port string) *MockSportsbookServer {
	return &MockSportsbookServer{BookName: name, Port: port}
}

func (m *MockSportsbookServer) StartMockStream() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range ticker.C {
			simulatedOdds := 1.85 + (rand.Float64() * 0.3)
			log.Printf("[Mock %s] Broadcasting live odds update: Decimal Odds = %.2f", m.BookName, simulatedOdds)
		}
	}()
}
