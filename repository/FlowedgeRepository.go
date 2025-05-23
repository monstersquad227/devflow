package repository

import (
	"devflow/model"
)

type FlowedgeRepository struct{}

func (f *FlowedgeRepository) ListFlowedges(pageNumber, pageSize int) ([]*model.Flowedge, error) {
	query := "SELECT agent_id, hostname, status, version, application, last_heartbeat, created_at, updated_at " +
		"FROM flowedge LIMIT ? OFFSET ?"
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Flowedge, 0)

	for rows.Next() {
		obj := &model.Flowedge{}
		if err = rows.Scan(&obj.AgentID, &obj.Hostname, &obj.Status, &obj.Version, &obj.Application, &obj.LastHeartBeat, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (f *FlowedgeRepository) CountFlowedges() (int, error) {
	var count int
	query := "SELECT COUNT(agent_id) " +
		"FROM flowedge"
	err := MysqlClient.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (f *FlowedgeRepository) GetFlowedgeByApplication(application string) (interface{}, error) {
	query := "SELECT agent_id, hostname " +
		"FROM flowedge " +
		"WHERE application = ? AND status = 'online'"
	rows, err := MysqlClient.Query(query, application)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Flowedge, 0)
	for rows.Next() {
		obj := &model.Flowedge{}
		if err = rows.Scan(&obj.AgentID, &obj.Hostname); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (f *FlowedgeRepository) CreateFlowedge(flowedge model.Flowedge) (int64, error) {
	query := "INSERT " +
		"INTO flowedge(agent_id, hostname, version, status) VALUES (?, ?, ?, ?) " +
		"ON DUPLICATE KEY UPDATE " +
		"hostname = VALUES(hostname), version = VALUES(version), status = VALUES(status)"
	result, err := MysqlClient.Exec(query, flowedge.AgentID, flowedge.Hostname, flowedge.Version, flowedge.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (f *FlowedgeRepository) UpdateFlowedgeLastHeartBeat(flow model.Flowedge) (int64, error) {
	query := "UPDATE " +
		"flowedge SET last_heartbeat = ? " +
		"WHERE agent_id = ?"
	result, err := MysqlClient.Exec(query, flow.LastHeartBeat, flow.AgentID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
