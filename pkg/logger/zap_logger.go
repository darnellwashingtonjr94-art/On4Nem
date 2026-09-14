package logger

import (
	"log"
)

type StructuredLogger struct {
	Environment string
}

func NewStructuredLogger(env string) *StructuredLogger {
	return &StructuredLogger{Environment: env}
}

func (sl *StructuredLogger) Info(msg string, fields map[string]interface{}) {
	log.Printf("[%s] INFO: %s | Fields: %v", sl.Environment, msg, fields)
}
