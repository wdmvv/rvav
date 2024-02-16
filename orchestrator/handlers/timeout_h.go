package handlers

import (
	"net/http"
	"orchestrator/config"
)

func TimeoutsHandler(w http.ResponseWriter, r *http.Request){
	WriteStruct(config.Conf.Signs, w, r)
}