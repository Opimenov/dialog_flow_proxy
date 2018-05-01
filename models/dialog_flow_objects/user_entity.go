//Contains a struct to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type UserEntity struct {
	SessionID string  `json:"sessionId,omitempty"`
	Name      string  `json:"name,omitempty"`
	Extend    bool    `json:"extend,omitempty"`
	Entries   []Entry `json:"entries,omitempty"`
}
