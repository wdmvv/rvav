package main

import (
	"orchestrator/config"
	"orchestrator/logging"
	"orchestrator/app"
)

func main() {
	logs.LoggerSetup()
	
	err := config.NewConfig("config/orchestrator.json")
	if err != nil {
		logs.ReportErr("failed to start config", err)
		return
	}

	app.StartServer(8000)
}
