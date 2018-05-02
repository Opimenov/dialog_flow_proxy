//Contains structs that define response objects
package responses

import (
	"time"
 ."leo/models/dialog_flow_objects"
)

//Contains fields that maps to dialog_flow response format
type AgentQueryResponse struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Result    Result    `json:"result,omitempty"`
	Status    Status    `json:"status,omitempty"`
	SessionID string    `json:"sessionId,omitempty"`
}

//Contains fields that maps to Result object response format from dialog_flow
type Result struct {
	Source           string              `json:"source,omitempty"`
	ResolvedQuery    string              `json:"resolvedQuery,omitempty"`
	Action           string              `json:"action,omitempty"`
	ActionIncomplete bool                `json:"actionIncomplete,omitempty"`
	Parameters       map[string]interface{} `json:"parameters,omitempty"`
	Contexts         []Context           `json:"contexts,omitempty"`
	Metadata         Metadata            `json:"metadata,omitempty"`

	Fulfillment      Fulfillment         `json:"fulfillment,omitempty"`
	Score            float32             `json:"score,omitempty"`
}

//Contains fields that are used for actually giving a response to user
//in a form of a single speech or an array of messages structs.
//Each message contains type of speech and a speech itself.
type Fulfillment struct {
	Speech   string    `json:"speech,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}
