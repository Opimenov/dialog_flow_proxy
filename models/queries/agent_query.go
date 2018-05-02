//Defines few  structs to model corresponding dialogflow objects.
package queries

import (
	."leo/models/dialog_flow_objects"
)

//Defines an agent query struct to interact with dialog flow api.
type AgentQuery struct {
	Query           string          `json:"query,omitempty"`
	E               Event           `json:"e,omitempty"`
	Event           Event           `json:"event,omitempty"`
	// V is a version of the protocol
	//https://dialogflow.com/docs/reference/agent/#protocol_version
	V               string          `json:"v,omitempty"`
	SessionID       string          `json:"sessionId,omitempty"`
	Lang            string          `json:"lang,omitempty"`
	Contexts        []Context       `json:"contexts,omitempty"`
	ResetContexts   bool            `json:"resetContexts,omitempty"`
	Entities        []Entity        `json:"entities,omitempty"`
	Timezone        string          `json:"timezone,omitempty"`
	Location        Location        `json:"location,omitempty"`
	OriginalRequest OriginalRequest `json:"originalRequest,omitempty"`
}

//Defines a location struct with two floats for location coordinates.
type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type OriginalRequest struct {
	Source string `json:"source,omitempty"`
	Data   string `json:"data,omitempty"`
}

