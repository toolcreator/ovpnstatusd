package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var labelNames = [...]string{"common_name"}

var (
	clientCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ovpnstatusd_client_count",
		Help: "The number of clients connected using the common name",
	}, labelNames[:])
)

func updateMetric(gaugeVec *prometheus.GaugeVec, commonName string, clientCount float64) {
	gauge, _ := gaugeVec.GetMetricWithLabelValues(commonName)
	gauge.Set(clientCount)
}

func resetMetrics() {
	clientCount.Reset()
}

func updateMetrics(ticker *time.Ticker, args *argBundle) {
	for {
		<-ticker.C
		values, err := getValues(args)
		if err != nil {
			log.Println(err)
		} else {
			resetMetrics()
			for _, value := range values {
				updateMetric(clientCount, value.commonName, float64(value.clientCount))
			}
		}
	}
}
