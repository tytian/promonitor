package main

import (
	"promonitor/server"
)

func init() {
	server.InitLog("debug")
}

func main() {
	server.InitDB()
	defer server.CloseDB()
	// run server
	server.StartMonitor(":8081")
}
