package repository

import "devflow/model"

type BuildTemplateRepository struct{}

/*
GetBuildTemplates 获取 build_template 记录
*/

func (b *BuildTemplateRepository) GetBuildTemplates(pageNumber, pageSize int) (interface{}, error) {
	query := "SELECT id, name, image_template_id, created_by, updated_by, created_at, updated_at " +
		"FROM build_template LIMIT ? OFFSET ? "
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.BuildTemplate, 0)
	for rows.Next() {
		obj := &model.BuildTemplate{}
		err = rows.Scan(&obj.ID, &obj.Name, &obj.ImageTemplateID, &obj.CreatedBy, &obj.UpdatedBy, &obj.CreatedAt,
			&obj.UpdatedAt)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetNameByID 通过 id 获取 build_template 记录的名称
*/

func (b *BuildTemplateRepository) GetNameByID(id int) (string, error) {
	var name string
	query := "SELECT name " +
		"FROM build_template " +
		"WHERE id = ? "
	err := MysqlClient.QueryRow(query, id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
