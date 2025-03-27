package repository

import "devflow/model"

type ImageRepository struct{}

func (i *ImageRepository) GetImages(pageNumber, pageSize int) ([]*model.Image, error) {
	query := "SELECT id, name, created_by, updated_by, created_at, updated_at " +
		"FROM image WHERE is_deleted = 0 LIMIT ? OFFSET ? "
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Image, 0)
	for rows.Next() {
		obj := &model.Image{}
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

func (i *ImageRepository) GetImagesCount() (int, error) {
	query := "SELECT count(id) " +
		"FROM image WHERE is_deleted = 0"
	var count int
	if err := MysqlClient.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (i *ImageRepository) GetImageNameById(id int) (string, error) {
	query := "SELECT name " +
		"FROM image WHERE id = ?"
	var imageName string
	if err := MysqlClient.QueryRow(query, id).Scan(&imageName); err != nil {
		return "", err
	}
	return imageName, nil
}
