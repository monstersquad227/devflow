package model

import "time"

type BuildTemplate struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	ImageTemplateID int       `json:"image_template_id,omitempty"`
	CreatedBy       string    `json:"created_by,omitempty"`
	UpdatedBy       string    `json:"updated_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
