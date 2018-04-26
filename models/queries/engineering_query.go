//Contains a struct to model particular type of query
package queries

import (
	//"fmt"
	//."leo/models/engineering_com_objects"
)

//Defines engineering.com query struct to interact with engineering.com api.
type EngineeringQuery struct {
	Query string `json:"query,omitempty"`
	Type string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	ProjectGoal string `json:"projectGoals,omitempty"`
	IsPrivate bool `json:"isPrivate,omitempty"`
	Description string `json:"description,omitempty"`
}
