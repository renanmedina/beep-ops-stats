package metrification

import "github.com/prometheus/client_golang/prometheus"

type CollectableMetrics struct {
	AttendanceDurationSeconds   prometheus.GaugeVec
	DisplacementDurationSeconds prometheus.GaugeVec
}

func NewCollectableMetrics(registry *prometheus.Registry) CollectableMetrics {
	collectable := CollectableMetrics{
		AttendanceDurationSeconds: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "attendance_duration_seconds",
			Help: "Attendance duration in seseconds",
		}, []string{"appointment_id", "order_ticket", "nurse_name", "operation_hub_name"}),
		DisplacementDurationSeconds: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "displacement_duration_seconds",
			Help:        "Displacement duration in seconds",
			ConstLabels: prometheus.Labels{},
		}, []string{"appointment_id", "order_ticket", "nurse_name", "operation_hub_name"}),
	}

	registry.MustRegister(collectable.AttendanceDurationSeconds)
	registry.MustRegister(collectable.DisplacementDurationSeconds)

	return collectable
}
