package execution

import (
	"fmt"
	"log"
)

type SlippageGuard struct {
	MaxAllowedSlippage float64
}

func NewSlippageGuard(maxSlippage float64) *SlippageGuard {
	return &SlippageGuard{
		MaxAllowedSlippage: maxSlippage,
	}
}

func (sg *SlippageGuard) ValidateSlippage(expectedPrice, actualPrice float64) error {
	if expectedPrice <= 0 {
		return fmt.Errorf("invalid expected price: %.2f", expectedPrice)
	}

	slippage := (actualPrice - expectedPrice) / expectedPrice
	if slippage < 0 {
		slippage = -slippage
	}

	if slippage > sg.MaxAllowedSlippage {
		log.Printf("Slippage threshold exceeded: %.4f > %.4f\n", slippage, sg.MaxAllowedSlippage)
		return fmt.Errorf("slippage %.4f exceeds maximum allowed threshold %.4f", slippage, sg.MaxAllowedSlippage)
	}

	return nil
}
