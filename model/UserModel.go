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

type Role struct {
	ID          int    `json:"id,omitempty"`
	RoleName    string `json:"role_name,omitempty"`
	RoleCode    string `json:"role_code,omitempty"`
	Description string `json:"description,omitempty"`
	Status      int    `json:"status,omitempty"`
	Deleted     int    `json:"deleted,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
type UserRole struct {
	ID        int    `json:"id,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	RoleID    int    `json:"role_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Permission struct {
	ID             int    `json:"id,omitempty"`
	PermissionName string `json:"permission_name,omitempty"`
	PermissionCode string `json:"permission_code,omitempty"`
	ResourceType   string `json:"resource_type,omitempty"`
	ParentID       int    `json:"parent_id,omitempty"`
	Path           string `json:"path,omitempty"`
	Description    string `json:"description,omitempty"`
	Status         int    `json:"status,omitempty"`
	Deleted        int    `json:"deleted,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

type RolePermission struct {
	ID           int    `json:"id,omitempty"`
	RoleID       int    `json:"role_id,omitempty"`
	PermissionID int    `json:"permission_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type Menu struct {
	ID             int    `json:"id,omitempty"`
	Path           string `json:"path,omitempty"`
	PermissionCode string `json:"permission_code,omitempty"`
	PermissionName string `json:"permission_name,omitempty"`
	//Children       []Menu `json:"children,omitempty"`
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token       string   `json:"token"`
	User        *User    `json:"user"`
	Roles       []*Role  `json:"roles"`
	Permissions []string `json:"permissions"`
	Menus       []*Menu  `json:"menus"`
}

type PasswordRequest struct {
	Account            string `json:"account" binding:"required"`
	Password           string `json:"password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}

type PasswordResponse struct {
	Message string `json:"message"`
}
