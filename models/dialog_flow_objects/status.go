//Contains a struct to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type Status struct {
	Code         int    `json:"code,omitempty"`
	ErrorDetails string `json:"errorDetails,omitempty"`
	ErrorID      string `json:"errorId,omitempty"`
	ErrorType    string `json:"errorType,omitempty"`
}
