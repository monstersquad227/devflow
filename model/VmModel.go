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
	ExpiredAt     string `json:"expired_at,omitempty"`
}

type VmCreateRequest struct {
	InstanceId    string `json:"instance_id"`
	InstanceName  string `json:"instance_name" binding:"required"`
	Password      string `json:"password" binding:"required"`
	PrivateIp     string `json:"private_ip" binding:"required"`
	PublicIp      string `json:"public_ip"`
	Spec          string `json:"spec" binding:"required"`
	Application   string `json:"application"`
	Region        string `json:"region" binding:"required"`
	CloudProvider string `json:"cloud_provider" binding:"required"`
	Os            string `json:"os" binding:"required"`
	ExpiredAt     string `json:"expired_at"`
}

type VmCreateResponse struct {
	LastInsertId int64 `json:"last_insert_id"`
}

type VmUpdateRequest struct {
	Id            int    `json:"_"`
	InstanceId    string `json:"instance_id" binding:"required"`
	InstanceName  string `json:"instance_name" binding:"required"`
	PrivateIp     string `json:"private_ip" binding:"required"`
	PublicIp      string `json:"public_ip"`
	Spec          string `json:"spec" binding:"required"`
	Application   string `json:"application"`
	Region        string `json:"region" binding:"required"`
	CloudProvider string `json:"cloud_provider" binding:"required"`
	Os            string `json:"os" binding:"required"`
}

type VmUpdateResponse struct {
	RowsAffected int64 `json:"rows_affected"`
}
