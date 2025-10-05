package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/renanmedina/beep-ops-stats/events"
)

func main() {
	registry := prometheus.NewRegistry()
	// Start generating events to parse into metrics
	go generateEvents(registry)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type EventPublishingStep struct {
	EventToSend         events.Event
	DelayToNextDuration time.Duration
}

type EventsDrivenScenario struct {
	MetricName     string
	steps          []EventPublishingStep
	elapsedTime    time.Duration
	metricRegister func(prometheusRegistry *prometheus.Registry, scenario *EventsDrivenScenario)
}

func (e *EventsDrivenScenario) registerMetric(prometheusRegistry *prometheus.Registry) {
	if e.metricRegister != nil {
		e.metricRegister(prometheusRegistry, e)
	}
}

func (e *EventsDrivenScenario) addElapsedTime(duration time.Duration) {
	e.elapsedTime += duration
}

func (e *EventsDrivenScenario) GetLabels() prometheus.Labels {
	labels := map[string]string{}
	for _, step := range e.steps {
		for key, value := range step.EventToSend.GetData() {
			labels[key] = value.(string)
		}
	}
	return labels
}

func generateEvents(prometheusRegistry *prometheus.Registry) {
	scenarios := []EventsDrivenScenario{
		{
			// Displacement duration metric
			MetricName: "displacement_duration_seconds",
			steps: []EventPublishingStep{
				{
					EventToSend: events.AppointmentJourneyEvent{
						JourneyStepName:  "displacement_started",
						AppointmentId:    "1",
						OrderTicket:      "1",
						NurseName:        "Técnica 1",
						OperationHubName: "São cristóvão",
						PublishedAt:      "2025-10-04T10:00:00Z",
					},
					DelayToNextDuration: 15 * time.Second,
				},
				{
					EventToSend: events.AppointmentJourneyEvent{
						JourneyStepName:  "arrived_in_place",
						AppointmentId:    "1",
						OrderTicket:      "1",
						NurseName:        "Técnica 1",
						OperationHubName: "São cristóvão",
						PublishedAt:      "2025-10-04T10:00:15Z",
					},
					DelayToNextDuration: 2 * time.Second,
				},
			},
			metricRegister: func(prometheusRegistry *prometheus.Registry, scenario *EventsDrivenScenario) {
				prometheusRegistry.MustRegister(
					prometheus.NewGaugeFunc(
						prometheus.GaugeOpts{
							Name:        scenario.MetricName,
							Help:        "The duration of the displacement",
							ConstLabels: scenario.GetLabels(),
						},
						func() float64 {
							return float64(scenario.elapsedTime.Seconds())
						},
					),
				)
			},
		},
		{
			// Attendance duration metric
			MetricName: "attendance_duration_seconds",
			steps: []EventPublishingStep{
				{
					EventToSend: events.AppointmentJourneyEvent{
						JourneyStepName:  "attendance_started",
						AppointmentId:    "1",
						OrderTicket:      "1",
						NurseName:        "Técnica 1",
						OperationHubName: "São cristóvão",
						PublishedAt:      "2025-10-04T10:00:17Z",
					},
					DelayToNextDuration: 10 * time.Second,
				},
				{
					EventToSend: events.AppointmentJourneyEvent{
						JourneyStepName:  "attendance_finished",
						AppointmentId:    "1",
						OrderTicket:      "1",
						NurseName:        "Técnica 1",
						OperationHubName: "São cristóvão",
						PublishedAt:      "2025-10-04T10:00:27Z",
					},
					DelayToNextDuration: 0,
				},
			},
			metricRegister: func(prometheusRegistry *prometheus.Registry, scenario *EventsDrivenScenario) {
				prometheusRegistry.MustRegister(
					prometheus.NewGaugeFunc(
						prometheus.GaugeOpts{
							Name:        scenario.MetricName,
							Help:        "The duration of the attendance",
							ConstLabels: scenario.GetLabels(),
						},
						func() float64 {
							return float64(scenario.elapsedTime.Seconds())
						},
					),
				)
			},
		},
	}

	for _, scenario := range scenarios {
		for _, step := range scenario.steps {
			fmt.Println("--------------------------------")
			fmt.Println(fmt.Sprintf("Scenario: %s", scenario.MetricName))
			fmt.Println(fmt.Sprintf("Sending event: %s - %s until next step event", step.EventToSend.GetName(), step.DelayToNextDuration.String()))
			scenario.addElapsedTime(step.DelayToNextDuration)
			time.Sleep(step.DelayToNextDuration)
		}
		fmt.Println(fmt.Sprintf("Scenario %s elapsed time: %s", scenario.MetricName, scenario.elapsedTime.String()))
		scenario.registerMetric(prometheusRegistry)
	}
}
