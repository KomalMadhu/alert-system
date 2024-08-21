package event

type Event struct {
	Client    string
	EventType string
	Timestamp int64
	Details   string
}
