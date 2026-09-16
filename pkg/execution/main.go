package execution

import (
	"fmt"
	"log"
)

type ExecutionEngine struct {
	Enabled bool
}

func NewExecutionEngine() *ExecutionEngine {
	return &ExecutionEngine{Enabled: true}
}

func (e *ExecutionEngine) ExecuteOrder(symbol string, amount float64) error {
	if !e.Enabled {
		return fmt.Errorf("execution engine is currently disabled")
	}
	log.Printf("Executing order for %s with amount %.2f\n", symbol, amount)
	return nil
}
