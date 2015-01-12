package main

import "encoding/json"

type EventPayload struct {
	Event string
	Time  int64
	Data  *json.RawMessage
}

func dispatch(e *EventPayload) error {
	switch e.Event {
	case "resource-merge":
		return handleResourceMerge(e)
	}

	return nil
}
