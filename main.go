package main

import (
	"promonitor/server"
)

func init() {
	server.InitDB()
}

func main() {
	defer server.CloseDB()
	// run server
	server.StartMonitor(":8081")
}
