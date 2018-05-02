//Contains structs to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object IntentAgent
type IntentAgent struct {
	ID             string       `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	ContextIn      []string     `json:"contextIn,omitempty"`
	ContextOut     []ContextOut `json:"contextOut,omitempty"`
	Actions        []string     `json:"actions,omitempty"`
	Parameters     []Parameter  `json:"parameters,omitempty"`
	Priority       int          `json:"priority,omitempty"`
	FallbackIntent bool         `json:"fallbackIntent,omitempty"`
}

//Dialog flow specific object
type ContextOut struct {
	Name     string `json:"name,omitempty"`
	Lifespan int    `json:"lifespan,omitempty"`
}
