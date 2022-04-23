package cli

import "quacc2/server/models"

func ListAgents() ([]models.Agent, error) {
	agents, err := models.ListAgents()
	//for _, agent := range agents {
	//	fmt.Printf("uuid: %v, ip: %v, hostname:%v, pid:%v, platform: %v\n",
	//		agent.AgentId,
	//		agent.Ips,
	//		agent.HostName,
	//		agent.Pid,
	//		agent.Platform,
	//	)
	//}
	return agents, err
}

func RunCommand(agentId, cmd string) error {
	c := models.NewCommand(agentId, cmd)
	has, err := models.ExistAgentId(agentId)
	if err != nil {
		return err
	}
	if has {
		err = c.Insert()
	}
	return err
}

func ListCommand(agentId string) ([]models.Command, error) {
	cmds, err := models.ListCommandByAgentId(agentId)
	if err != nil {
		return cmds, err
	}

	//for _, cmd := range cmds {
	//	fmt.Printf("agent: %v, cmd: %v, status: %v, time: %v\n", cmd.AgentId, cmd.Content, cmd.Status, cmd.CreateTime)
	//}

	return cmds, err
}