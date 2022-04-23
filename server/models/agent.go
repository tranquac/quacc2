package models

import "time"

type Agent struct {
	Id           int64
	AgentId      string    `json:"agent_id"`
	Platform     string    `json:"platform"`
	Architecture string    `json:"architecture"`
	UserName     string    `json:"user_name"`
	UserGUID     string    `json:"user_guid"`
	HostName     string    `json:"host_name"`
	Ips          []string  `json:"ips" xorm:"text"`
	Pid          int       `json:"pid"`
	Debug        bool      `json:"debug"`
	Proto        string    `json:"proto"`
	UserAgent    string    `json:"user_agent"`
	Initial      bool      `json:"initial"`
	CreateTime   time.Time `xorm:"created"`
	UpdateTime   time.Time `xorm:"updated"`
	Version      int       `xorm:"version"`
}

//Insert a agent to DB
func (a *Agent) Insert() error {
	_, err := Engine.Insert(a)
	return err
}

//List agent have in DB
func ListAgents() ([]Agent, error) {
	agents := make([]Agent, 0)
	err := Engine.Find(&agents)
	return agents, err
}

//Update agent info to DB
func UpdateAgent(agentId string) error {
	agent := new(Agent)
	has, err := Engine.Where("agent_id=?", agentId).Get(agent)
	if err != nil {
		return err
	}
	if has {
		_, err = Engine.Id(agent.Id).Update(agent)
	}
	return err
}

//Check agent have in DB
func ExistAgentId(agentId string) (bool, error) {
	agent := new(Agent)
	has, err := Engine.Where("agent_id=?", agentId).Get(agent)
	return has, err
}

//Remove agent in DB
func RemoveAll() error {
	_, err := Engine.Exec("delete from agent")
	return err
}