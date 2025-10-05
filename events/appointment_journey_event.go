package events

type AppointmentJourneyEvent struct {
	JourneyStepName  string `json:"journey_step_name"`
	AppointmentId    string `json:"appointment_id"`
	OrderTicket      string `json:"order_ticket"`
	NurseName        string `json:"nurse_name"`
	OperationHubName string `json:"operation_hub_name"`
	PublishedAt      string `json:"published_at"`
}

func (e AppointmentJourneyEvent) GetName() string {
	return e.JourneyStepName
}

func (e AppointmentJourneyEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"appointment_id":     e.AppointmentId,
		"order_ticket":       e.OrderTicket,
		"nurse_name":         e.NurseName,
		"operation_hub_name": e.OperationHubName,
		"published_at":       e.PublishedAt,
	}
}
