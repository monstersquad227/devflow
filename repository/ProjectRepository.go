package repository

import (
	"devflow/model"
	"errors"
)

type ProjectRepository struct{}

func (r *ProjectRepository) ListProjects(pageNumber, pageSize int) ([]*model.Project, error) {
	query := "SELECT id, gitlab_name, deployment_name, gitlab_id, gitlab_repo, task_id, " +
		"project_build_path, project_package_name, description " +
		"FROM project WHERE is_deleted = 0 LIMIT ? OFFSET ? "

	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.Project, 0)
	for rows.Next() {
		obj := &model.Project{}
		err = rows.Scan(&obj.ID, &obj.GitlabName, &obj.DeploymentName, &obj.GitlabID, &obj.GitlabRepo,
			&obj.TaskID, &obj.ProjectBuildPath, &obj.ProjectPackageName, &obj.Description)
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

func (r *ProjectRepository) CountProjects() (int, error) {
	query := "SELECT count(id) " +
		"FROM project WHERE is_deleted = 0"
	var count int
	err := MysqlClient.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProjectRepository) CreateProject(project *model.Project) (int64, error) {
	query := "INSERT " +
		"INTO project(gitlab_name, deployment_name, task_id, gitlab_id, gitlab_repo, project_build_path, project_package_name, description) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, project.GitlabName, project.DeploymentName, project.TaskID, project.GitlabID, project.GitlabRepo,
		project.ProjectBuildPath, project.ProjectPackageName, project.Description)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if lastId < 0 {
		return 0, errors.New("插入失败")
	}
	return lastId, nil
}

func (r *ProjectRepository) ExistDeploymentName(name string) bool {
	var total int
	query := "SELECT COUNT(deployment_name) " +
		"FROM project " +
		"WHERE deployment_name = ?"
	err := MysqlClient.QueryRow(query, name).Scan(&total)
	if err != nil {
		return false
	}
	if total > 0 {
		return true
	}
	return false
}

func (r *ProjectRepository) UpdateProject(project *model.Project) (int64, error) {
	query := "UPDATE project " +
		"SET deployment_name = ?, gitlab_id = ?, gitlab_repo = ?, task_id = ?, project_build_path = ?, project_package_name = ? " +
		"WHERE id = ?"
	result, err := MysqlClient.Exec(query, project.DeploymentName, project.GitlabID, project.GitlabRepo, project.TaskID, project.ProjectBuildPath, project.ProjectPackageName, project.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *ProjectRepository) DeleteProject(id int) (int64, error) {
	query := "UPDATE project " +
		"set is_deleted = 1 WHERE id = ?"
	result, err := MysqlClient.Exec(query, id)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, err
}

func (r *ProjectRepository) ListProjectApplications() ([]*model.Project, error) {
	query := "SELECT deployment_name " +
		"FROM project "
	rows, err := MysqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.Project, 0)
	for rows.Next() {
		obj := &model.Project{}
		err = rows.Scan(&obj.DeploymentName)
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

func (r *ProjectRepository) GetIdByDeploymentName(name string) (int, error) {
	var id int
	query := "SELECT id " +
		"FROM project WHERE deployment_name = ?"
	if err := MysqlClient.QueryRow(query, name).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProjectRepository) GetBuildTemplateIDByID(id int) (int, error) {
	var buildTemplateID int
	query := "SELECT build_template_id " +
		"FROM project " +
		"WHERE id = ?"
	err := MysqlClient.QueryRow(query, id).Scan(&buildTemplateID)
	if err != nil {
		return 0, err
	}
	return buildTemplateID, nil
}
