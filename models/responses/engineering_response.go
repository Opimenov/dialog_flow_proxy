//Contains EngineeringQueryResponse struct that defines
//an engineering.com api response object.
package responses

import (
	"time"
	."leo/models/dialog_flow_objects"
)

//Contains fields that maps to engineering.com response format:
// ID, Timestamp, Lang, Result, Status, SessionID
type EngineeringQueryResponse struct {
	//TODO. Since we didn't use Engineering.com api we don't know the
	//TODO exact response format. So this may need to be modified.
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Result    Result    `json:"result,omitempty"`
	Status    Status    `json:"status,omitempty"`
	SessionID string    `json:"sessionId,omitempty"`
} 