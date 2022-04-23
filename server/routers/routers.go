package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quacc2/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var agent models.Agent
	err := c.BindJSON(&agent)
	fmt.Println(agent, err)
	agentId := agent.AgentId
	has, err := models.ExistAgentId(agentId)
	if err == nil && has {
		models.UpdateAgent(agentId)
	} else {
		err = agent.Insert()
		fmt.Println(err)
	}
}

//send command to agent
func GetCommand(c *gin.Context) {
	agnetId := c.Param("uuid")
	cmds, _ := models.ListCommandByAgentId(agnetId)
	cmdJson, _ := json.Marshal(cmds)
	fmt.Println(agnetId, string(cmdJson))
	c.JSON(http.StatusOK, cmds)
}

//receive result from client
func SendResult(c *gin.Context) {
	cmdId := c.Param("id")
	result := c.PostForm("result")
	id, _ := strconv.Atoi(cmdId)
	err := models.UpdateCommandResult(int64(id), result)
	fmt.Println(cmdId, result, err, c.Request.PostForm)
	if err == nil {
		err = models.SetCmdStatusToFinished(int64(id))
	}
}