//Contains a struct to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type Metadata struct {
	IntentID                  string `json:"intentId,omitempty"`
	WebhookUsed               string `json:"webhookUsed,omitempty"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed,omitempty"`
	IntentName                string `json:"intentName,omitempty"`
}
