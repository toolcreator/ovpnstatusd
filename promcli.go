package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var labelNames = [...]string{"example_label"}

var (
	exampleGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ovpnstatusdexample_gauge",
		Help: "Example Gauge",
	}, labelNames[:])
)

func updateMetric(gaugeVec *prometheus.GaugeVec, example_label string, value float64) {
	gauge, _ := gaugeVec.GetMetricWithLabelValues(example_label)
	gauge.Set(value)
}

func resetMetrics() {
	exampleGauge.Reset()
}

func updateMetrics(ticker *time.Ticker) {
	for {
		<-ticker.C
		values, err := getValues()
		if err != nil {
			log.Println(err)
		} else {
			resetMetrics()
			for _, value := range values {
				updateMetric(exampleGauge, value.example_label, float64(value.value))
			}
		}
	}
}
