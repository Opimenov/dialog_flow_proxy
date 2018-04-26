//Contains structs that map to json strings received as responses
package responses

import (
	"time"
	."leo/models/dialog_flow_objects"
)

//Contains fields that maps to engineering.com response format
type EngineeringQueryResponse struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Result    Result    `json:"result,omitempty"`
	Status    Status    `json:"status,omitempty"`
	SessionID string    `json:"sessionId,omitempty"`
} 