package service

import (
	"devflow/model"
	pb "github.com/monstersquad227/flowedge-proto"
	"io"
	"log"
	"sync"
	"time"
)

type FlowEdgeServer struct {
	pb.UnimplementedFlowEdgeServer
	streams         sync.Map
	pendingResponse sync.Map
}

var GlobalFlowEdgeServer = &FlowEdgeServer{}

func (s *FlowEdgeServer) Communicate(stream pb.FlowEdge_CommunicateServer) error {
	var agentID string
	client := NewFlowedgeService()
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Agent %s disconnected", agentID)
			if agentID != "" {
				s.streams.Delete(agentID)
			}
			return nil
		}
		if err != nil {
			log.Printf("Recv error from agent %s: %v", agentID, err)
			if agentID != "" {
				s.streams.Delete(agentID)
			}
			return err
		}

		switch msg.Type {
		/*
			注册 agent
		*/
		case pb.MessageType_REGISTER:
			agentID = msg.GetRegister().AgentId
			s.streams.Store(agentID, stream)

			f := model.Flowedge{
				AgentID:  agentID,
				Hostname: msg.GetRegister().Hostname,
				Version:  msg.GetRegister().Version,
				Status:   "online",
			}
			result, err := client.Create(f)
			if err != nil {
				return err
			}
			log.Printf("insert %d: ", result)
			//log.Printf("Register: %+v", msg.GetRegister())
			log.Printf("Agent %s registered", agentID)

		/*
			agent 心跳
		*/
		case pb.MessageType_HEARTBEAT:
			f := model.Flowedge{
				AgentID:       msg.GetHeartbeat().AgentId,
				LastHeartBeat: time.Now().Format("2006-01-02 15:04:05"),
			}
			_, err = client.Update(f)
			if err != nil {
				return err
			}
			//fmt.Printf("Agent %s updated, %d", agentID, result)
			//log.Printf("Heartbeat from %s", msg.GetHeartbeat().AgentId)

		/*
			执行指令
		*/
		case pb.MessageType_EXECUTE_RESPONSE:
			r := msg.GetExecuteResponse()
			log.Printf("Execution result: %s, output: %s, error: %s", r.CommandId, r.Output, r.Error)

			if chVal, ok := s.pendingResponse.Load(r.CommandId); ok {
				ch := chVal.(chan *pb.ExecuteResponse)
				ch <- r
				s.pendingResponse.Delete(r.CommandId)
			}
		}
	}
}
