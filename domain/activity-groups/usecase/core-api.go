package usecase

import (
	"golang-todo/entity"
	"golang-todo/valueobject"
	"time"
)

func (activityGroups activityGroupsUsecase) GetAll(param map[string]interface{}) (response []valueobject.ActivityGroups, err error) {
	response, err = activityGroups.mysqlRepository.GetAll(param)
	return
}

func (activityGroups activityGroupsUsecase) GetOne(param map[string]interface{}) (response valueobject.ActivityGroups, err error) {
	response, err = activityGroups.mysqlRepository.GetOne(param)
	return
}

func (activityGroups activityGroupsUsecase) Store(payload valueobject.ActivityGroups) (valueobject.ActivityGroups, error) {
	now := time.Now()
	var storeData valueobject.ActivityGroupsPayloadInsert

	payload.ID, _ = activityGroups.mysqlRepository.GenerateID()
	payload.CreatedAt = &now
	payload.UpdatedAt = &now
	storeData.Data = append(storeData.Data, payload)

	queryConfig, err := activityGroups.ProcessStore(storeData)

	if err != nil {
		return payload, err
	}

	return payload, activityGroups.mysqlRepository.Exec(queryConfig...)
}

func (activityGroups activityGroupsUsecase) Update(payload valueobject.ActivityGroups) error {
	now := time.Now()
	var updateData valueobject.ActivityGroupsPayloadUpdate

	updateData.Data = append(updateData.Data, valueobject.ActivityGroupsDataUpdate{
		Param: valueobject.ActivityGroups{
			StandardKey: entity.StandardKey{
				ID: payload.ID,
			},
		},
		Body: valueobject.ActivityGroups{
			ActivityGroups: entity.ActivityGroups{
				Title: payload.Title,
				Email: payload.Email,
			},
			Time: entity.Time{
				UpdatedAt: &now,
			},
		},
	})

	queryConfig, err := activityGroups.ProcessUpdate(updateData)

	if err != nil {
		return err
	}

	return activityGroups.mysqlRepository.Exec(queryConfig...)
}

func (activityGroups activityGroupsUsecase) Delete(payload valueobject.ActivityGroups) (err error) {
	var deleteData valueobject.ActivityGroupsPayloadDelete
	deleteData.Param = append(deleteData.Param, valueobject.ActivityGroups{
		StandardKey: entity.StandardKey{
			ID: payload.ID,
		},
	})
	queryConfig, err := activityGroups.ProcessDelete(deleteData)

	if err != nil {
		return
	}

	return activityGroups.mysqlRepository.Exec(queryConfig...)
}
