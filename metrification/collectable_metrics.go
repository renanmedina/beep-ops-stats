package metrification

import "github.com/prometheus/client_golang/prometheus"

type CollectableMetrics struct {
	AttendanceDurationSeconds   prometheus.Gauge
	DisplacementDurationSeconds prometheus.Gauge
}

func NewCollectableMetrics(registry *prometheus.Registry) CollectableMetrics {
	return CollectableMetrics{
		AttendanceDurationSeconds: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "attendance_duration_seconds",
			Help: "Attendance duration in seseconds",
		}),
		DisplacementDurationSeconds: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "displacement_duration_seconds",
			Help: "Displacement duration in seconds",
		}),
	}
}
