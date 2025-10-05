package scenarios

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type EventDrivenScenario struct {
	MetricName     string
	Steps          []EventStep
	elapsedTime    time.Duration
	MetricRegister func(prometheusRegistry *prometheus.Registry, scenario *EventDrivenScenario)
}

func (e *EventDrivenScenario) RegisterMetric(prometheusRegistry *prometheus.Registry) {
	if e.MetricRegister != nil {
		e.MetricRegister(prometheusRegistry, e)
	}
}

func (e *EventDrivenScenario) AddElapsedTime(duration time.Duration) {
	e.elapsedTime += duration
}

func (e *EventDrivenScenario) GetElapsedTime() time.Duration {
	return e.elapsedTime
}

func (e *EventDrivenScenario) GetLabels() prometheus.Labels {
	labels := map[string]string{}
	for _, step := range e.Steps {
		for key, value := range step.EventReceived.GetData() {
			labels[key] = value.(string)
		}
	}
	return labels
}
