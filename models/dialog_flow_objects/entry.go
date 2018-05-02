package dialog_flow_objects

//Dialog flow specific object
type Entry struct {
	Value    string   `json:"value,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
}
