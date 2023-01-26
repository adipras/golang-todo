package entity

type TodoItemsKey struct {
	ActivityGroupID uint64 `json:"activity_group_id"`
}

type TodoItems struct {
	TodoItemsKey
	Title    string `json:"title,omitempty" binding:"required"`
	IsActive bool   `json:"is_active,omitempty"`
	Priority string `json:"priority,omitempty"`
}
