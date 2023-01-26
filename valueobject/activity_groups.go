package valueobject

import "golang-todo/entity"

type ActivityGroups struct {
	entity.ActivityGroups
	entity.StandardKey
	entity.Time
}

type ActivityGroupsPayloadInsert struct {
	Data []ActivityGroups `json:"data" binding:"required"`
	User string
}

type ActivityGroupsPayloadUpdate struct {
	Data []ActivityGroupsDataUpdate `json:"data" binding:"required"`
	User string
}

type ActivityGroupsDataUpdate struct {
	Param ActivityGroups `json:"param" binding:"required"`
	Body  ActivityGroups `json:"body" binding:"required"`
}

type ActivityGroupsPayloadDelete struct {
	Param []ActivityGroups `json:"param" binding:"required"`
}
