package scenarios

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/renanmedina/beep-ops-stats/metrification"
)

type EventDrivenScenario struct {
	Name            string
	Steps           []EventStep
	elapsedTime     time.Duration
	MetricsRegister func(collectableMetrics *metrification.CollectableMetrics, scenario *EventDrivenScenario)
}

func (e *EventDrivenScenario) RegisterMetrics(collectableMetrics *metrification.CollectableMetrics) {
	if e.MetricsRegister != nil {
		e.MetricsRegister(collectableMetrics, e)
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
