package model

type ProjectBuild struct {
	Id          int    `json:"id,omitempty"`
	ProjectId   int    `json:"project_id,omitempty"`
	JenkinsId   int    `json:"jenkins_id,omitempty"`
	TaskName    string `json:"task_name,omitempty"`
	BuildStatus string `json:"build_status,omitempty"`
	BuildParams string `json:"build_params,omitempty"`
	CreateBy    string `json:"create_by,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	UpdateTime  string `json:"update_time,omitempty"`
}
