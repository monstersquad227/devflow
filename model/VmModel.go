package model

type Vm struct {
	Id            int    `json:"id,omitempty"`
	InstanceId    string `json:"instance_id,omitempty"`
	InstanceName  string `json:"instance_name,omitempty"`
	PrivateIp     string `json:"private_ip,omitempty"`
	PublicIp      string `json:"public_ip,omitempty"`
	Spec          string `json:"spec,omitempty"`
	Application   string `json:"application,omitempty"`
	Region        string `json:"region,omitempty"`
	CloudProvider string `json:"cloud_provider,omitempty"`
	Os            string `json:"os,omitempty"`
	Password      string `json:"password,omitempty"`
	IsDeleted     int    `json:"is_deleted,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}
