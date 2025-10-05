package scenarios

import (
	"math/rand"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/renanmedina/beep-ops-stats/events"
	"github.com/renanmedina/beep-ops-stats/metrification"
)

var scenariosFactories = map[string]func(faker *faker.Faker) EventDrivenScenario{
	"displacement_duration_seconds": func(faker *faker.Faker) EventDrivenScenario {
		appointmentId := faker.UUID().V4()
		orderTicket := faker.UUID().V4()
		nurseName := faker.Person().Name()
		operationHubName := faker.Lorem().Word()
		publishedAt := time.Now()
		delayToNextDuration := time.Duration(faker.Int32Between(15, 60)) * time.Second

		return EventDrivenScenario{
			// Displacement duration metric
			Name: "Displacement duration seconds",
			Steps: []EventStep{
				{
					EventReceived: events.AppointmentJourneyEvent{
						JourneyStepName:  "displacement_started",
						AppointmentId:    appointmentId,
						OrderTicket:      orderTicket,
						NurseName:        nurseName,
						OperationHubName: operationHubName,
						PublishedAt:      publishedAt.Format(time.RFC3339),
					},
					DelayToNextDuration: delayToNextDuration,
				},
				{
					EventReceived: events.AppointmentJourneyEvent{
						JourneyStepName:  "arrived_in_place",
						AppointmentId:    appointmentId,
						OrderTicket:      orderTicket,
						NurseName:        nurseName,
						OperationHubName: operationHubName,
						PublishedAt:      publishedAt.Add(delayToNextDuration).Format(time.RFC3339),
					},
					DelayToNextDuration: 0,
				},
			},
			MetricsRegister: func(collectableMetrics *metrification.CollectableMetrics, scenario *EventDrivenScenario) {
				collectableMetrics.DisplacementDurationSeconds.Set(scenario.GetElapsedTime().Seconds())
			},
		}
	},
	"attendance_duration_seconds": func(faker *faker.Faker) EventDrivenScenario {
		appointmentId := faker.UUID().V4()
		orderTicket := faker.UUID().V4()
		nurseName := faker.Person().Name()
		operationHubName := faker.Lorem().Word()
		publishedAt := time.Now()
		delayToNextDuration := time.Duration(faker.Int32Between(5, 60)) * time.Second

		return EventDrivenScenario{
			// Attendance duration metric
			Name: "Attendance duration seconds",
			Steps: []EventStep{
				{
					EventReceived: events.AppointmentJourneyEvent{
						JourneyStepName:  "attendance_started",
						AppointmentId:    appointmentId,
						OrderTicket:      orderTicket,
						NurseName:        nurseName,
						OperationHubName: operationHubName,
						PublishedAt:      publishedAt.Format(time.RFC3339),
					},
					DelayToNextDuration: delayToNextDuration,
				},
				{
					EventReceived: events.AppointmentJourneyEvent{
						JourneyStepName:  "attendance_finished",
						AppointmentId:    appointmentId,
						OrderTicket:      orderTicket,
						NurseName:        nurseName,
						OperationHubName: operationHubName,
						PublishedAt:      publishedAt.Add(delayToNextDuration).Format(time.RFC3339),
					},
					DelayToNextDuration: 0,
				},
			},
			MetricsRegister: func(collectableMetrics *metrification.CollectableMetrics, scenario *EventDrivenScenario) {
				collectableMetrics.AttendanceDurationSeconds.Set(float64(scenario.GetElapsedTime().Seconds()))
			},
		}
	},
}

func NewScenario(scenarioName string) EventDrivenScenario {
	scenarioFactory := scenariosFactories[scenarioName]
	faker := faker.New()
	return scenarioFactory(&faker)
}

func getAvailableScenarios() ([]string, int) {
	availableScenarios := make([]string, 0, len(scenariosFactories))
	for key := range scenariosFactories {
		availableScenarios = append(availableScenarios, key)
	}
	return availableScenarios, len(availableScenarios)
}

func Generate(quantity int) []EventDrivenScenario {
	scenarios := make([]EventDrivenScenario, quantity)
	availableScenarios, scenariosCount := getAvailableScenarios()
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < quantity; i++ {
		randomIndex := randomizer.Intn(scenariosCount)
		scenarioName := availableScenarios[randomIndex]
		scenarios[i] = NewScenario(scenarioName)
	}
	return scenarios
}
