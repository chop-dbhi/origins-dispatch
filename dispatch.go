package main

import "encoding/json"

type EventPayload struct {
	Event string
	Time  int64
	Data  *json.RawMessage
}

func dispatch(e *EventPayload) {
	switch e.Event {
	case "resource-merge":
		handleResourceMerge(e)
		break
	}
}
