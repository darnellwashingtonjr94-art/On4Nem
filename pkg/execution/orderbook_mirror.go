package execution

import (
	"sync"
)

type OrderBookMirror struct {
	mu     sync.RWMutex
	bids   map[string]float64
	asks   map[string]float64
}

func NewOrderBookMirror() *OrderBookMirror {
	return &OrderBookMirror{
		bids: make(map[string]float64),
		asks: make(map[string]float64),
	}
}

func (obm *OrderBookMirror) UpdateBook(marketID string, bid float64, ask float64) {
	obm.mu.Lock()
	defer obm.mu.Unlock()
	obm.bids[marketID] = bid
	obm.asks[marketID] = ask
}
