package rollingwindows

type Event int8

const EventCount = 4

const (
	EVENT_SUCCESS	Event = 0
	EVENT_FAILURE	Event = 1
	EVENT_TIMEOUT	Event = 2
	EVENT_REJECTION	Event = 3
)
