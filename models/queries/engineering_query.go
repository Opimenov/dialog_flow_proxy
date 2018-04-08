package queries

import (
	//"fmt"
	//."leo/models/engineering_com_objects"
)

type EngineeringQuery struct {
	Query string `json:"query,omitempty"`
	Type string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	ProjectGoal string `json:"projectGoals,omitempty"`
	IsPrivate bool `json:"isPrivate,omitempty"`
	Description string `json:"description,omitempty"`
}
