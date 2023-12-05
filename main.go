package main

import (
	"promonitor/monitor"
	"promonitor/server"
)

func init() {
	server.InitLog("debug")
}

func main() {
	monitor.InitDB()
	defer monitor.CloseDB()
	// run server
	server.StartMonitor(":8081")
}
