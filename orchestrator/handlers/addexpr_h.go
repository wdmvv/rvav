package handlers

//for adding expression to the job list
import (
	"encoding/json"
	"net/http"
	"orchestrator/calc"
	"strings"
)

func AddExprHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data AddExprReqIn
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := AddExprReqOut{0, "failed to calculate expression"}
		WriteStruct(e, w, r)
		return
	}
	
	data.sanitize()
	post, err := calc.InfixToPostfix(data.Expr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := AddExprReqOut{0, err.Error()}
		WriteStruct(e, w, r)
		return
	}

	res, err := calc.Eval(post)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		e := AddExprReqOut{0, err.Error()}
		WriteStruct(e, w, r)
		return
	}
	e := AddExprReqOut{res, ""}
	WriteStruct(e, w, r)
}

// to remove spaces in expr
func (d *AddExprReqIn) sanitize(){
	d.Expr = strings.ReplaceAll(d.Expr, " ", "")
}