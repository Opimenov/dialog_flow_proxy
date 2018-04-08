package responses

import (
	"time"
	."leo/models/dialog_flow_objects"
)
//TODO find out what we are getting in response from engineering.com and
// create appropriate struct

type EngineeringQueryResponse struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Result    Result    `json:"result,omitempty"`
	Status    Status    `json:"status,omitempty"`
	SessionID string    `json:"sessionId,omitempty"`
} 