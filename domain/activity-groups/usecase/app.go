package usecase

import (
	activityGroups "golang-todo/domain/activity-groups"
)

/*
*
the struct of usecases
*/
type activityGroupsUsecase struct {
	mysqlRepository activityGroups.MysqlRepository
}

/*
*
the initiator function for usecase
*/
func NewActivityGroupsUsecase(mysql activityGroups.MysqlRepository) activityGroups.Usecase {
	return &activityGroupsUsecase{
		mysqlRepository: mysql,
	}
}
