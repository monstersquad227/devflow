package model

import "time"

type Project struct {
	ID                 uint      `json:"id,omitempty"`
	GitlabName         string    `json:"gitlab_name,omitempty"`
	DeploymentName     string    `json:"deployment_name,omitempty"`
	GitlabID           int       `json:"gitlab_id,omitempty"`
	GitlabRepo         string    `json:"gitlab_repo,omitempty"`
	BuildTemplateID    *uint     `json:"build_template_id,omitempty"`
	ProjectBuildPath   string    `json:"project_build_path,omitempty"`
	ProjectPackageName string    `json:"project_package_name,omitempty"`
	Description        string    `json:"description,omitempty"`
	IsDeleted          bool      `json:"is_deleted,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}
