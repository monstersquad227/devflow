package repository

import "devflow/model"

type TaskRepository struct{}

func (t *TaskRepository) GetTasks(pageNumber, pageSize int) ([]*model.Task, error) {
	query := "SELECT id, name, image_id, created_by, updated_by, created_at, updated_at " +
		"FROM task WHERE is_deleted = 0 LIMIT ? OFFSET ?"
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Task, 0)
	for rows.Next() {
		obj := &model.Task{}
		if err = rows.Scan(&obj.Id, &obj.Name, &obj.ImageID, &obj.CreatedBy, &obj.UpdatedBy, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (t *TaskRepository) GetTasksCount() (int, error) {
	query := "SELECT count(id) " +
		"FROM task WHERE is_deleted = 0 "
	var count int
	if err := MysqlClient.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (t *TaskRepository) GetTaskNameANDImageIDById(id int) (string, int, error) {
	query := "SELECT name, image_id " +
		"FROM task WHERE id = ?"
	var name string
	var imageId int
	if err := MysqlClient.QueryRow(query, id).Scan(&name, &imageId); err != nil {
		return "", 0, err
	}
	return name, imageId, nil
}
