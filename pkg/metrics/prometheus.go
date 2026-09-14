package metrics

import "log"

type PrometheusExporter struct {
	Port string
}

func NewPrometheusExporter(port string) *PrometheusExporter {
	return &PrometheusExporter{Port: port}
}

func (pe *PrometheusExporter) StartExporter() {
	log.Printf("Starting Prometheus telemetry exporter on port %s...", pe.Port)
}
