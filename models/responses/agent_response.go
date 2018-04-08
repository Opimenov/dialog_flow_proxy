package responses

import (
	"time"
 ."leo/models/dialog_flow_objects"
)

type AgentQueryResponse struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Result    Result    `json:"result,omitempty"`
	Status    Status    `json:"status,omitempty"`
	SessionID string    `json:"sessionId,omitempty"`
}

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

type Fulfillment struct {
	Speech   string    `json:"speech,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}
