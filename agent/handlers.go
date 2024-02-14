package main

//for handlers
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Agent is running!")
}

func evalHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data evalReqIn
	err := decoder.Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := evalReqOut{0, fmt.Sprintf("failed to parse incoming json - %s", err)}
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
	time.Sleep(time.Duration(e.Timeout) * time.Millisecond)
	return res, nil
}
