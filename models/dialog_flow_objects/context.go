//Contains structs to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type Context struct {
	Name       string           `json:"name,omitempty"`
	Lifespan   int              `json:"lifespan,omitempty"`
	Parameters ContextParameter `json:"parameters,omitempty"`
}

//Dialog flow specific object
type ContextParameter struct {
	IntentAction string `json:"intent_action,omitempty"`
	Name         string `json:"name,omitempty"`
	Value        string `json:"value,omitempty"`
}
