package metrics

import "log"

type CustomCollector struct {
	MetricName string
}

func (cc *CustomCollector) RegisterCounter(name string) {
	log.Printf("Registering custom Prometheus telemetry counter: %s", name)
}
