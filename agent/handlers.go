package main

//for handlers
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math"

	"agent/logging"
)

// all in & out requests structs

type evalReqIn struct {
	Op1     float64 `json:"op1"`
	Op2     float64 `json:"op2"`
	Sign    string  `json:"sign"`
	Timeout int     `json:"timeout"`
}

type evalReqOut struct {
	Result float64 `json:"result"`
	Errmsg string  `json:"errmsg"`
}

type statReqOut struct {
	Msg string `json:"msg"`
}

// they end here

// handlers for endpoints

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.ReportAction(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	})
}

// /status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	e := statReqOut{"agent is running!"}
	msg, _ := json.Marshal(e)
	w.Write(msg)
}

// /eval
func EvalHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data evalReqIn
	err := decoder.Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := evalReqOut{0, "invalid request body"}
		msg, _ := json.Marshal(e)
		w.Write(msg)
		return
	}

	Limit.Acquire(context.Background(), 1)
	defer Limit.Release(1)

	res, err := data.Eval()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := evalReqOut{res, fmt.Sprintf("failed to calculate expression - %s", err)}

		msg, _ := json.Marshal(e)
		w.Write(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
	e := evalReqOut{res, ""}
	msg, _ := json.Marshal(e)
	w.Write(msg)
}

func (e *evalReqIn) Eval() (float64, error) {
	var res float64
	if e.Sign == "+" {
		res = e.Op1 + e.Op2
	} else if e.Sign == "-" {
		res = e.Op1 - e.Op2
	} else if e.Sign == "*" {
		res = e.Op1 * e.Op2
	} else if e.Sign == "/" {
		if e.Op2 == 0 {
			return 0, fmt.Errorf("division by zero detected")
		}
		res = e.Op1 / e.Op2
	} else {
		return 0, fmt.Errorf("invalid operator detected")
	}
	//precision moment, imagine better round function
	res = math.Round(res * 100) / 100
	time.Sleep(time.Duration(e.Timeout) * time.Millisecond)
	return res, nil
}
