//Contains a struct to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type Parameter struct {
	Name         string   `json:"name,omitempty"`
	Value        string   `json:"value,omitempty"`
	DefaultValue string   `json:"defaultValue,omitempty"`
	Required     bool     `json:"required,omitempty"`
	DataType     string   `json:"dataType,omitempty"`
	Prompts      []string `json:"prompts,omitempty"`
	IsList       bool     `json:"isList,omitempty"`
}
