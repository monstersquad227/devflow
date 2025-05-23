package model

type ProjectDeploy struct {
	PublishType string   `json:"publish_type,omitempty"`
	Name        string   `json:"name,omitempty"`
	Env         string   `json:"env,omitempty"`
	Namespace   string   `json:"namespace,omitempty"`
	Ecs         []string `json:"ecs,omitempty"`
	Tag         string   `json:"tag,omitempty"`
}
