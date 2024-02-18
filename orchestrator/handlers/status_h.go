package handlers

import (
	"net/http"
)

// /status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	e := StatusReqOut{"orchestrator is running!"}
	WriteStruct(e, w, r)
}