//Contains a struct to model particular type of query
package queries

import (
	"fmt"
	"reflect"
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

//defines a location struct with two floats for location coordinates
type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type OriginalRequest struct {
	Source string `json:"source,omitempty"`
	Data   string `json:"data,omitempty"`
}


func (query AgentQuery) ToMap() map[string]string {
	params := make(map[string]string)

	if query.Query != "" {
		params["query"] = query.Query
	}

	if !reflect.DeepEqual(query.E, Event{}) {
		params["e"] = query.Event.Name
	}

	if !reflect.DeepEqual(query.Contexts, []Context{}) && len(query.Contexts) > 0 {
		params["contexts"] = query.Contexts[0].Name
	}

	if !reflect.DeepEqual(query.Location, Location{}) {
		params["latitude"] = fmt.Sprintf("%f", query.Location.Latitude)
		params["longitude"] = fmt.Sprintf("%f", query.Location.Longitude)
	}

	params["v"] = query.V
	params["sessionId"] = query.SessionID
	params["lang"] = query.Lang

	return params
}
