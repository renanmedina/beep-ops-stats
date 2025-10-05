package events

type Event interface {
	GetName() string
	GetData() map[string]interface{}
}
