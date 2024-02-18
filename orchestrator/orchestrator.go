package main

import (
	"orchestrator/app"
	"orchestrator/config"
	"orchestrator/db"
	"orchestrator/logging"
)

func main() {
	logging.LoggerSetup()

	err := config.NewConfig("config/orchestrator.json")
	if err != nil {
		logging.ReportErr("failed to start config", err)
		return
	}
	if config.Conf.UseDB{
		err = db.GetConn(config.Conf.User, config.Conf.Pswd, config.Conf.DBname, config.Conf.TabName)
		if err != nil {
			logging.ReportErr("error on db init", err)
		}

		err = db.DBCOnn.Table()
		if err != nil {
			logging.ReportErr("failed to connect to db", err)
	}}
	app.StartServer(8000)
}
