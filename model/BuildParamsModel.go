package model

//type BuildParamsV2 struct {
//	Project             string `json:"project,omitempty"`
//	DependentProject    string `json:"dependent_project,omitempty"`
//	AliasName           string `json:"alias_name,omitempty"`
//	Branch              string `json:"branch,omitempty"`
//	DependentBranch     string `json:"dependent_branch,omitempty"`
//	Repository          string `json:"repository,omitempty"`
//	DependentRepository string `json:"dependent_repository,omitempty"`
//	EnvironmentUnique   string `json:"environment_unique,omitempty"`
//	BuildPath           string `json:"build_path,omitempty"`
//	PackageName         string `json:"package_name,omitempty"`
//	ImageSource         string `json:"image_source,omitempty"`
//	Command             string `json:"command,omitempty"`
//	CreateBy            string `json:"create_by,omitempty"`
//	ShortID             string `json:"short_id,omitempty"`
//}

type BuildParams struct {
	GitlabName        string `json:"gitlab_name,omitempty"`
	DeploymentName    string `json:"deployment_name,omitempty"`
	TaskID            string `json:"task_id,omitempty"`
	Branch            string `json:"branch,omitempty"`
	GitlabRepo        string `json:"gitlab_repo,omitempty"`
	ImageSource       string `json:"image_source,omitempty"`
	EnvironmentUnique string `json:"environment_unique,omitempty"`
	HarborUrl         string `json:"harbor_url,omitempty"`
	ShortId           string `json:"short_id,omitempty"`
	Command           string `json:"command,omitempty"`
	CreatedBy         string `json:"created_by,omitempty"`
	DotnetBuildParams
}

type DotnetBuildParams struct {
	DependGitlabRepo   string `json:"depend_gitlab_repo,omitempty"`
	DependBranch       string `json:"depend_branch,omitempty"`
	ProjectBuildPath   string `json:"project_build_path,omitempty"`
	ProjectPackageName string `json:"project_package_name,omitempty"`
}
