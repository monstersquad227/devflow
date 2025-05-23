package repository

import "devflow/model"

type EnvRepository struct{}

func (e *EnvRepository) ListEnvs(pageNumber, pageSize int) ([]*model.Env, error) {
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

func (e *EnvRepository) CountEnvs() (int, error) {
	query := "SELECT count(id) " +
		"FROM env WHERE is_deleted = 0 "
	var count int
	if err := MysqlClient.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (e *EnvRepository) CreateEnv(env *model.Env) (int64, error) {
	query := "INSERT " +
		"INTO env(name, created_by, updated_by) VALUES (?, ?, ?)"
	result, err := MysqlClient.Exec(query, env.Name, env.CreatedBy, env.UpdatedBy)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (e *EnvRepository) DeleteEnv(id int) (int64, error) {
	query := "UPDATE env " +
		"SET is_deleted = 1 " +
		"WHERE id = ?"
	result, err := MysqlClient.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (e *EnvRepository) UpdateEnv(env *model.Env) (int64, error) {
	query := "UPDATE env " +
		"SET name = ?, updated_by = ? " +
		"WHERE id = ?"
	result, err := MysqlClient.Exec(query, env.Name, env.UpdatedBy, env.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
