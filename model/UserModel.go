package model

type User struct {
	ID          int      `json:"id,omitempty"`
	Account     string   `json:"account,omitempty"`
	Password    string   `json:"password,omitempty"`
	Name        string   `json:"name,omitempty"`
	Email       string   `json:"email,omitempty"`
	Mobile      string   `json:"mobile,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Deleted     int      `json:"deleted,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty"`
	UpdatedAt   string   `json:"updated_at,omitempty"`
}
