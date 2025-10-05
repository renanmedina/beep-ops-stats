package picking

type PickingFinishedEvent struct {
	PickingId        int32  `json:"picking_id"`
	NurseId          string `json:"nurse_id"`
	OperationHubName string `json:"operation_hub_name"`
	FinishedAt       string `json:"finished_at"`
}

func (e PickingFinishedEvent) GetName() string {
	return "picking_finished"
}

func (e PickingFinishedEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"picking_id":         e.PickingId,
		"nurse_id":           e.NurseId,
		"operation_hub_name": e.OperationHubName,
		"finished_at":        e.FinishedAt,
	}
}
