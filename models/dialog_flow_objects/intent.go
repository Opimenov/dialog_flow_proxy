//Contains a struct to model particular type of dialog flow object
package dialog_flow_objects

//Dialog flow specific object
type Intent struct {
	ID                    string         `json:"id,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Auto                  bool           `json:"auto,omitempty"`
	Contexts              []string       `json:"contexts,omitempty"`
	Templates             []string       `json:"templates,omitempty"`
	UserSays              []UserSay      `json:"userSays,omitempty"`
	Responses             []Response     `json:"responses,omitempty"`
	Priority              int            `json:"priority,omitempty"`
	WebhookUsed           bool           `json:"webhookUsed,omitempty"`
	WebhookForSlotFilling bool           `json:"webhookForSlotFilling,omitempty"`
	FallbackIntent        bool           `json:"fallbackIntent,omitempty"`
	CortanaCommand        CortanaCommand `json:"cortanaCommand,omitempty"`
	Events                []Event        `json:"events,omitempty"`
}

//Dialog flow specific object
type CortanaCommand struct {
	NavigationOrService string `json:"navigationOrService,omitempty"`
	Target              string `json:"target,omitempty"`
}
//Dialog flow specific object
type UserSay struct {
	ID         string `json:"userSays,omitempty"`
	Data       Data   `json:"data,omitempty"`
	IsTemplate bool   `json:"isTemplate,omitempty"`
	Count      int    `json:"count,omitempty"`
}
//Dialog flow specific object
type Response struct {
	Action           string      `json:"action,omitempty"`
	ResetContexts    bool        `json:"resetContexts,omitempty"`
	AffectedContexts []Context   `json:"affectedContexts,omitempty"`
	Parameters       []Parameter `json:"parameters,omitempty"`
	Messages         []Message   `json:"message,omitempty"`
}
//Dialog flow specific object
type Data struct {
	Text        string `json:"text,omitempty"`
	Meta        string `json:"meta,omitempty"`
	Alias       string `json:"alias,omitempty"`
	UserDefined bool   `json:"userDefined,omitempty"`
}
