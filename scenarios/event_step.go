package scenarios

import (
	"time"

	"github.com/renanmedina/beep-ops-stats/events"
)

type EventStep struct {
	EventReceived       events.Event
	DelayToNextDuration time.Duration
}
