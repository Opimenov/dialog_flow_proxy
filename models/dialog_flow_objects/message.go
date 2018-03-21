package dialog_flow_objects

type Message struct {
	Type   int    `json:"type,omitempty"`
	Speech string `json:"speech,omitempty"`
}
