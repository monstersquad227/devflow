package model

type ProjectBuild struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"project_id"`
	JenkinsId   int    `json:"jenkins_id"`
	BuildStatus string `json:"build_status"`
	BuildParams string `json:"build_params"`
	CreateBy    string `json:"create_by"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
