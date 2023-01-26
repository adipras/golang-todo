package entity

import "time"

type StandardKey struct {
	ID   uint64 `json:"id"`
	UUID string `json:"uuid,omitempty"`
}

type Time struct {
	CreatedBy string     `json:"createdBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedBy string     `json:"updateBy,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
