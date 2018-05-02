//Contains structs to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type Event struct {
	Name string            `json:"name,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}
