package repository

import (
	"golang-todo/utils/database"
	"golang-todo/valueobject"
)

func (db *mysqlTodoItemsRepository) GetAll(param map[string]interface{}) (response []valueobject.TodoItems, err error) {
	var result valueobject.TodoItems

	builder := database.New(MYSQL, MYSQL_TABLE, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"todo_id", "activity_group_id", "title", "is_active", "priority", "created_at", "updated_at"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query, err := db.sqlDB.Query(builder.Result.Query, builder.Result.Value...)

	if err != nil {
		return
	}

	defer query.Close()

	for query.Next() {
		err = query.Scan(
			&result.ID,
			&result.ActivityGroupID,
			&result.Title,
			&result.IsActive,
			&result.Priority,
			&result.CreatedAt,
			&result.UpdatedAt,
		)

		if err != nil {
			return
		}

		response = append(response, result)
	}

	return
}

func (db *mysqlTodoItemsRepository) GetOne(param map[string]interface{}) (response valueobject.TodoItems, err error) {
	builder := database.New(MYSQL, MYSQL_TABLE, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"todo_id", "activity_group_id", "title", "is_active", "priority", "created_at", "updated_at"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query := db.sqlDB.QueryRow(builder.Result.Query, builder.Result.Value...)

	err = query.Scan(
		&response.ID,
		&response.ActivityGroupID,
		&response.Title,
		&response.IsActive,
		&response.Priority,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return
}

func (db *mysqlTodoItemsRepository) Store(column []string, data []interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(MYSQL, MYSQL_TABLE, INSERT)

	builder.OnInsert = database.OnInsert{
		Column: column,
		Data:   data,
	}

	err = builder.QueryBuilder()
	return
}

func (db *mysqlTodoItemsRepository) Update(param map[string]interface{}, data map[string]interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(MYSQL, MYSQL_TABLE, UPDATE)

	builder.OnUpdate = database.OnUpdate{
		Where: param,
		Data:  data,
	}

	err = builder.QueryBuilder()
	return
}

func (db *mysqlTodoItemsRepository) Delete(param map[string]interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(MYSQL, MYSQL_TABLE, DELETE)

	builder.OnDelete = database.OnDelete{
		Where: param,
	}

	err = builder.QueryBuilder()
	return
}
