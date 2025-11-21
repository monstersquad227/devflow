package model

type Flowedge struct {
	AgentID       string  `json:"agent_id,omitempty"`
	Hostname      string  `json:"hostname,omitempty"`
	Status        string  `json:"status,omitempty"`
	Version       string  `json:"version,omitempty"`
	Application   *string `json:"application,omitempty"`
	LastHeartBeat string  `json:"last_heartbeat,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
}

type FlowedgePatchRequest struct {
	AgentID     string `json:"agent_id"`
	Application string `json:"application" binding:"required"`
}

type FlowedgePatchResponse struct {
	RowsAffected int64 `json:"rows_affected"`
}
