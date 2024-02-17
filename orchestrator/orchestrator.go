package main

import (
	"orchestrator/app"
	"orchestrator/config"
	"orchestrator/db"
	"orchestrator/logging"
)

func main() {
	logs.LoggerSetup()

	err := config.NewConfig("config/orchestrator.json")
	if err != nil {
		logs.ReportErr("failed to start config", err)
		return
	}
	err = db.GetConn(config.Conf.User, config.Conf.Pswd, config.Conf.DBname, config.Conf.TabName)
	if err != nil{
		logs.ReportErr("error on db init", err)
	}

	err = db.DBCOnn.Table()
	if err != nil{
		logs.ReportErr("failed to connect to db", err)
	}
	
	app.StartServer(8000)
}
