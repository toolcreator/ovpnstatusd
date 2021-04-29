package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type argBundle struct {
	remotePath  *string
	destination *string
	user        *string
	password    *string
	timeout     *uint
}

func main() {
	port := flag.Int("port", 8080, "The port ovpnstatusd listens on")
	updateInterval := flag.Int("interval", 60000, "The update interval in milliseconds")
	remotePath := flag.String("remote-path", "/etc/openvpn/openvpn-status.log",
		"The path to openvpn-status.log at the destination")
	destination := flag.String("destination", "",
		"The hostname/IP address and port of the destination, separated by colon.")
	user := flag.String("user", "", "The username")
	password := flag.String("password", "", "The password")
	timeout := flag.Uint("timeout", 5, "The timeout in seconds")
	flag.Parse()

	statsUpdateTicker := time.NewTicker(time.Duration(*updateInterval) * time.Millisecond)
	go updateMetrics(statsUpdateTicker,
		&argBundle{
			remotePath:  remotePath,
			destination: destination,
			user:        user,
			password:    password,
			timeout:     timeout,
		})

	http.Handle("/metrics", promhttp.Handler())
	err := make(chan error)
	go func() {
		err <- http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	}()

	log.Println(fmt.Sprintf("Listening on port %d.", *port))
	log.Println("Metrics are exposed at /metrics endpoint.")

	<-err
	log.Println(err)
}
