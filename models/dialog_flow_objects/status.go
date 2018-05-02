//Contains structs to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object Status
//https://dialogflow.com/docs/reference/agent/
type Status struct {
	Code         int    `json:"code,omitempty"`
	ErrorDetails string `json:"errorDetails,omitempty"`
	ErrorID      string `json:"errorId,omitempty"`
	ErrorType    string `json:"errorType,omitempty"`
}
