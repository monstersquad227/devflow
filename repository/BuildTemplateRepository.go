package repository

//
//import "devflow/model"
//
//type BuildTemplateRepository struct{}
//
///*
//GetBuildTemplates 获取 build_template 记录
//*/
//
//func (b *BuildTemplateRepository) GetBuildTemplates(pageNumber, pageSize int) (interface{}, error) {
//	query := "SELECT id, name, image_template_id, created_by, updated_by, created_at, updated_at " +
//		"FROM build_template LIMIT ? OFFSET ? "
//	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	data := make([]*model.BuildTemplate, 0)
//	for rows.Next() {
//		obj := &model.BuildTemplate{}
//		err = rows.Scan(&obj.ID, &obj.Name, &obj.ImageTemplateID, &obj.CreatedBy, &obj.UpdatedBy, &obj.CreatedAt,
//			&obj.UpdatedAt)
//		if err != nil {
//			return nil, err
//		}
//		data = append(data, obj)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
///*
//GetNameByID 通过 id 获取 build_template 记录的名称
//*/
//
//func (b *BuildTemplateRepository) GetNameByID(id int) (string, int, error)
//	var buildTemplateName string
//	var imageTemplateId int
//	query := "SELECT name, image_template_id " +
//		"FROM build_template " +
//		"WHERE id = ? "
//	err := MysqlClient.QueryRow(query, id).Scan(&buildTemplateName, &imageTemplateId)
//	if err != nil {
//		return "", 0, err
//	}
//	return buildTemplateName, imageTemplateId, nil
//}
