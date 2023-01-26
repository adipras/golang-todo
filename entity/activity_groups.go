package entity

type ActivityGroups struct {
	Title string `json:"title,omitempty" binding:"required"`
	Email string `json:"email,omitempty"`
}
