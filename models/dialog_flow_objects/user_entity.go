package dialog_flow_objects

//Dialog flow specific object UserEntity
//https://dialogflow.com/docs/entities
type UserEntity struct {
	SessionID string  `json:"sessionId,omitempty"`
	Name      string  `json:"name,omitempty"`
	Extend    bool    `json:"extend,omitempty"`
	Entries   []Entry `json:"entries,omitempty"`
}
