package main

import (
	"quacc2/client/util"
)

func main() {
	go util.Ping()
	util.Command()
}