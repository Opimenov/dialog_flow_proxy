package engineering_com_objects

type Project struct {
	Type string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	ProjectGoal string `json:"projectGoals,omitempty"`
	IsPrivate bool `json:"isPrivate,omitempty"`
	Description string `json:"description,omitempty"`
}
