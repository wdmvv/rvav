package main

import (
	"orchestrator/config"
	"orchestrator/logging"
)

func main(){
	err := config.NewConfig("config/orchestrator.json")
	if err != nil{
		logs.ReportErr("failed to start config", err)
		return
	}
	
}