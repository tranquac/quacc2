package cli

import (
	"errors"
	"quacc2/server/models"
	"strings"

	"github.com/abiosoft/ishell"
)

func Shell() error {
	var err error
	shell := ishell.New()
	shell.Println("Command & Control manager")

	// list agent
	shell.AddCmd(&ishell.Cmd{
		Name: "list",
		Help: "list agent",
		Func: func(c *ishell.Context) {
			agents, err := ListAgents()
			if err == nil {
				DisplayAgent(agents)
			}
		},
	})

	// list command
	shell.AddCmd(&ishell.Cmd{
		Name: "cmd",
		Func: func(c *ishell.Context) {
			if len(c.Args) == 0 {
				c.Err(errors.New("Missing agent_id"))
			} else {
				agentId := c.Args[0]
				cmds, err := ListCommand(agentId)
				if err == nil {
					DisplayCommand(cmds)
				}
			}
		},
		Help: "List Command",
		Completer: func(args []string) []string {
			agentList := make([]string, 0)
			agents, err := ListAgents()
			if err == nil {
				for _, agent := range agents {
					agentList = append(agentList, agent.AgentId)
				}
			}

			return agentList
		},
	})

	// add command
	shell.AddCmd(&ishell.Cmd{
		Name: "run",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 2 {
				c.Err(errors.New("Missing agent_id"))
			} else {
				agentId := c.Args[0]
				cmd := c.Args[1:]
				c := strings.Join(cmd, " ")
				_ = RunCommand(agentId, c)
			}
		},
		Help: "run agent_id command",
		Completer: func(args []string) []string {
			agentList := make([]string, 0)
			agents, err := ListAgents()
			if err == nil {
				for _, agent := range agents {
					agentList = append(agentList, agent.AgentId)
				}
			}
			return agentList
		},
	})

	// remove all agents
	shell.AddCmd(&ishell.Cmd{
		Name: "remove",
		Func: func(c *ishell.Context) {
			_ = models.RemoveAll()
		},
		Help: "remove all agent",
	})

	go ListCmdResult()

	shell.Run()

	return err
}