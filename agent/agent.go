package main

// main load balancer
import (
	logs "agent/logging"
	"agent/workers"
)

func main() {
	logs.LoggerSetup()
	workers.Set()

	StartServer(8001)
}
