package handlers

import (
	"net/http"
	"encoding/json"
)

// /status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	e := StatusReqOut{"agent is running!"}
	msg, _ := json.Marshal(e)
	w.Write(msg)
}