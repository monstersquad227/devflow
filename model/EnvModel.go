package model

type Env struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Remark    string `json:"remark,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	IsDeleted int    `json:"is_deleted,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type EnvCreateRequest struct {
	Name      string `json:"name" binding:"required,min=2,max=50"`
	Remark    string `json:"remark"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type EnvCreateResponse struct {
	LastInsertId int64 `json:"last_insert_id"`
}

type EnvUpdateRequest struct {
	Id        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Remark    string `json:"remark"`
	UpdatedBy string `json:"updated_by"`
}

type EnvUpdateResponse struct {
}
