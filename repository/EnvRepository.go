package repository

import "devflow/model"

type EnvRepository struct{}

func (e *EnvRepository) GetEnvs(pageNumber, pageSize int) ([]*model.Env, error) {
	query := "SELECT id, name, created_by, updated_by, created_at, updated_at " +
		"FROM env WHERE is_deleted = 0 LIMIT ? OFFSET ?"
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Env, 0)

	for rows.Next() {
		obj := &model.Env{}
		if err = rows.Scan(&obj.Id, &obj.Name, &obj.CreatedBy, &obj.UpdatedBy, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (e *EnvRepository) GetEnvsCount() (int, error) {
	query := "SELECT count(id) " +
		"FROM env WHERE is_deleted = 0 "
	var count int
	if err := MysqlClient.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
