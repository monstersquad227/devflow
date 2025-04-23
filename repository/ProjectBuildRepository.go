package repository

import (
	"devflow/model"
)

type ProjectBuildRepository struct {
}

/*
GetProjectIdByStatus 获取 ING 状态的项目
*/

func (p *ProjectBuildRepository) GetProjectIdByStatus() ([]int, error) {
	query := "SELECT project_id " +
		"FROM project_build WHERE build_status = 'ING'"
	rows, err := MysqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	data := make([]int, 0)
	for rows.Next() {
		obj := model.ProjectBuild{}
		err := rows.Scan(&obj.ProjectId)
		if err != nil {
			return nil, err
		}
		data = append(data, obj.ProjectId)
	}
	return data, nil
}

/*
CreateProjectBuild 创建 project_build 记录
*/

func (p *ProjectBuildRepository) CreateProjectBuild(params, createBy, taskName string, projectID int, jenkinsID int64) (int64, error) {
	query := "INSERT " +
		"INTO project_build (project_id, jenkins_id, task_name, build_status, build_params, create_by) " +
		"VALUES(?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, projectID, jenkinsID, taskName, "ING", params, createBy)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

/*
UpdateBuildStatus 通过 job_id project_id 来更新构建状态
*/

func (p *ProjectBuildRepository) UpdateBuildStatus(projectId, jenkinsId int, status string) (int64, error) {
	query := "UPDATE project_build " +
		"SET build_status = ? " +
		"WHERE project_id = ? AND jenkins_id = ? "

	result, err := MysqlClient.Exec(query, status, projectId, jenkinsId)
	if err != nil {
		return 0, err
	}
	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAff, nil
}

/*
GetProjectBuildByProjectId 通过 project_id 获取 project_build 记录
*/

func (p *ProjectBuildRepository) GetProjectBuildByProjectId(projectID int) (interface{}, error) {
	query := "SELECT id, jenkins_id, task_name, build_status, build_params, create_by, create_time, update_time " +
		"FROM project_build " +
		"WHERE project_id = ? ORDER BY id DESC"
	rows, err := MysqlClient.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	data := make([]*model.ProjectBuild, 0)
	for rows.Next() {
		obj := model.ProjectBuild{}
		err = rows.Scan(&obj.Id, &obj.JenkinsId, &obj.TaskName, &obj.BuildStatus, &obj.BuildParams, &obj.CreateBy, &obj.CreateTime, &obj.UpdateTime)
		if err != nil {
			return nil, err
		}
		data = append(data, &obj)
	}
	return data, nil
}

func (p *ProjectBuildRepository) GetProjectBuildCountByProjectId(projectId int) (int, error) {
	query := "SELECT COUNT(id) " +
		"FROM project_build WHERE project_id = ?"
	var count int
	err := MysqlClient.QueryRow(query, projectId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
