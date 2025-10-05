package picking

type PickingStartedEvent struct {
	PickingId        int32  `json:"picking_id"`
	NurseId          string `json:"nurse_id"`
	OperationHubName string `json:"operation_hub_name"`
	StartedAt        string `json:"started_at"`
}

func (e PickingStartedEvent) GetName() string {
	return "picking_started"
}

func (e PickingStartedEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"picking_id":         e.PickingId,
		"nurse_id":           e.NurseId,
		"operation_hub_name": e.OperationHubName,
		"started_at":         e.StartedAt,
	}
}
