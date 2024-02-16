package handlers

import (
	"encoding/json"
	"net/http"
	"orchestrator/config"
)

func TimeoutsHandler(w http.ResponseWriter, r *http.Request){
	msg, _ := json.Marshal(config.Conf.Signs)
	w.Write(msg)
}