package main

import (
	"fmt"
	"net/http"
	"time"

	"log/slog"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/renanmedina/beep-ops-stats/metrification"
	"github.com/renanmedina/beep-ops-stats/scenarios"
)

func main() {
	registry := prometheus.NewRegistry()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	collectableMetrics := metrification.NewCollectableMetrics(registry)
	// Start generating events to parse into metrics
	go runRandomScenarios(&collectableMetrics, logger)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	logger.Info("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	logger.Info("Server started on port 8080")
}

func runRandomScenarios(collectableMetrics *metrification.CollectableMetrics, logger *slog.Logger) {
	scenarios := scenarios.Generate(2)

	for {
		logger.Info("Running events receiving scenarios")
		for _, scenario := range scenarios {
			logger.Info(fmt.Sprintf("Scenario: %s", scenario.Name))
			for _, step := range scenario.Steps {
				logger.Info(fmt.Sprintf("Step: %s", step.EventReceived.GetName()))
				logger.Info(fmt.Sprintf("Receiving event: %s - waiting %s until next step event", step.EventReceived.GetName(), step.DelayToNextDuration.String()))
				scenario.AddElapsedTime(step.DelayToNextDuration)
				time.Sleep(step.DelayToNextDuration)
			}
			logger.Info(fmt.Sprintf("Scenario %s elapsed time: %s", scenario.Name, scenario.GetElapsedTime().String()))
			scenario.RegisterMetrics(collectableMetrics)
		}
		logger.Info("Sleeping for 1 minutes")
		time.Sleep(1 * time.Minute)
	}
}
