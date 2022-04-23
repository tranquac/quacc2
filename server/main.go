package main

import (
	"fmt"
	"os"
	"quacc2/server/cli"
	"quacc2/server/models"
	"quacc2/server/routers"
	"strings"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/ping", routers.Ping)
	r.POST("/cmd/:uuid", routers.GetCommand)
	r.POST("/send_result/:id", routers.SendResult)

	return r
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%v [remove_agent|list_agent|list_cmd|run command|serv|shell]\n", os.Args[0])
		os.Exit(0)
	}
	cmd := strings.ToLower(os.Args[1])
	parameters := ""
	if len(os.Args) > 2 {
		parameters = strings.Join(os.Args[2:], " ")
	}
	switch cmd {
	case "serv":
		_ = models.RemoveAll()
		r := setupRouter()
		//err := r.Run(":8080")
		err := r.RunTLS(":8080", "./certs/server.pem", "./certs/server.key")
		_ = err
	case "run":
		fmt.Printf("run %v", parameters)
		if len(os.Args) >= 3 {
			agent := os.Args[2]
			c := strings.Join(os.Args[3:], " ")
			err := cli.RunCommand(agent, c)
			_ = err
		}
	case "list_agent":
		_, _ = cli.ListAgents()
	case "list_cmd":
		_, _ = cli.ListCommand(parameters)
	case "remove_agent":
		_ = models.RemoveAll()
	case "shell":
		_ = cli.Shell()
	}

}