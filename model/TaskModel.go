package model

type Task struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ImageID   int    `json:"image_id,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	IsDeleted int    `json:"is_deleted,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
