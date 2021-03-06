// +build !notime

package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/log"
)

type timeCollector struct {
	metric prometheus.Counter
}

func init() {
	Factories["time"] = NewTimeCollector
}

// Takes a prometheus registry and returns a new Collector exposing
// the current system time in seconds since epoch.
func NewTimeCollector() (Collector, error) {
	return &timeCollector{
		metric: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "time",
			Help:      "System time in seconds since epoch (1970).",
		}),
	}, nil
}

func (c *timeCollector) Update(ch chan<- prometheus.Metric) (err error) {
	now := float64(time.Now().Unix())
	log.Debugf("Set time: %f", now)
	c.metric.Set(now)
	c.metric.Collect(ch)
	return err
}
