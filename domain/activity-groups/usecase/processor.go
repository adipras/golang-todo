package usecase

import (
	"golang-todo/utils/database"
	"golang-todo/valueobject"
)

func (activityGroups activityGroupsUsecase) ProcessStore(payload valueobject.ActivityGroupsPayloadInsert) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		/**
		add data you wanted to insert on this interface{}...
		*/
		data := []interface{}{
			[]interface{}{
				x.ID,
				x.Title,
				x.Email,
				x.CreatedAt,
				x.UpdatedAt,
			},
		}

		/**
		column on data and this line should have same order
		*/
		column := []string{
			"activity_id", "title", "email", "created_at", "updated_at",
		}

		queryInsert, err := activityGroups.mysqlRepository.Store(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (activityGroups activityGroupsUsecase) ProcessUpdate(payload valueobject.ActivityGroupsPayloadUpdate) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"activity_id": x.Param.ID, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"title":      x.Body.Title, // add the data to update the row
			"email":      x.Body.Email,
			"updated_at": x.Body.UpdatedAt,
		}

		queryUpdate, err := activityGroups.mysqlRepository.Update(param, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (activityGroups activityGroupsUsecase) ProcessDelete(payload valueobject.ActivityGroupsPayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"activity_id": x.ID, // add the parameter to delete the row
			},
		}

		queryDelete, err := activityGroups.mysqlRepository.Delete(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
