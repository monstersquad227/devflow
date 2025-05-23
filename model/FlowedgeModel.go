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
