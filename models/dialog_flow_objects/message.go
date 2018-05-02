//Contains structs to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object Message
type Message struct {
	Type   int    `json:"type,omitempty"`
	Speech string `json:"speech,omitempty"`
}
