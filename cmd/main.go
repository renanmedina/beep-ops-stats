package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/renanmedina/beep-ops-stats/events"
	"github.com/renanmedina/beep-ops-stats/scenarios"
)

func main() {
	registry := prometheus.NewRegistry()
	// Start generating events to parse into metrics
	go generateEvents(registry)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateEvents(prometheusRegistry *prometheus.Registry) {
	scenarios := []scenarios.EventDrivenScenario{
		{
			// Displacement duration metric
			MetricName: "displacement_duration_seconds",
			Steps: []scenarios.EventStep{
				{
					EventReceived: events.AppointmentJourneyEvent{
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
					EventReceived: events.AppointmentJourneyEvent{
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
			MetricRegister: func(prometheusRegistry *prometheus.Registry, scenario *scenarios.EventDrivenScenario) {
				prometheusRegistry.MustRegister(
					prometheus.NewGaugeFunc(
						prometheus.GaugeOpts{
							Name:        scenario.MetricName,
							Help:        "The duration of the displacement",
							ConstLabels: scenario.GetLabels(),
						},
						func() float64 {
							return float64(scenario.GetElapsedTime().Seconds())
						},
					),
				)
			},
		},
		{
			// Attendance duration metric
			MetricName: "attendance_duration_seconds",
			Steps: []scenarios.EventStep{
				{
					EventReceived: events.AppointmentJourneyEvent{
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
					EventReceived: events.AppointmentJourneyEvent{
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
			MetricRegister: func(prometheusRegistry *prometheus.Registry, scenario *scenarios.EventDrivenScenario) {
				prometheusRegistry.MustRegister(
					prometheus.NewGaugeFunc(
						prometheus.GaugeOpts{
							Name:        scenario.MetricName,
							Help:        "The duration of the attendance",
							ConstLabels: scenario.GetLabels(),
						},
						func() float64 {
							return float64(scenario.GetElapsedTime().Seconds())
						},
					),
				)
			},
		},
	}

	for _, scenario := range scenarios {
		for _, step := range scenario.Steps {
			fmt.Println("--------------------------------")
			fmt.Println(fmt.Sprintf("Scenario: %s", scenario.MetricName))
			fmt.Println(fmt.Sprintf("Sending event: %s - %s until next step event", step.EventReceived.GetName(), step.DelayToNextDuration.String()))
			scenario.AddElapsedTime(step.DelayToNextDuration)
			time.Sleep(step.DelayToNextDuration)
		}
		fmt.Println(fmt.Sprintf("Scenario %s elapsed time: %s", scenario.MetricName, scenario.GetElapsedTime().String()))
		scenario.RegisterMetric(prometheusRegistry)
	}
}
