// metrics/metrics.go
package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (

	// Response time histogram
	ResponseTimeHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "url_shortener_time_seconds",
		Help:    "Histogram of response times for requests.",
		Buckets: prometheus.DefBuckets,
	})

	// Uptime gauge
	UptimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "url_shortener_uptime_seconds",
		Help: "Uptime of the application in seconds.",
	})

	// Error rate counter
	ErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "url_shortener_errors_total",
		Help: "Total number of errors.",
	})

	// Throughput counter
	ThroughputCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "url_shortener_requests_total",
		Help: "Total number of requests.",
	})
)

func init() {
	// Register your metrics
	prometheus.MustRegister(ResponseTimeHistogram, UptimeGauge, ErrorCounter, ThroughputCounter)

}
func StartUptimeTracker() {
	go func() {
		start := time.Now()
		for {
			UptimeGauge.Set(time.Since(start).Seconds())
			time.Sleep(1 * time.Second)
		}
	}()
}
