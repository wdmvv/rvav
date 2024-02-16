package handlers

// for all request structs to be jsoned
import (
	"encoding/json"
	"net/http"
)

// /chtime

type ChtimeReqIn struct{
	Sign string `json:"sign"`
	Ms int `json:"ms"`
}

type ChtimeReqOut struct{
	Error string `json:"error"`
}

// /status

type TimeoutsReqOut struct{
	Plus int `json:"plus"`
	Minus int `json:"minus"`
	Mul int `json:"mul"`
	Div int `json:"div"`
}

// /timeouts
type StatusReqOut struct{
	Msg string `json:"msg"`
}

// /addexpr
type AddExprReqIn struct{
	Expr string `json:"expr"`
}

type AddExprReqOut struct{
	Result float64 `json:"result"`
	Errmsg string `json:"errmsg"`
}

// because i got tired of doing this manually every time i need to report error/anything
// converts structure into json representation
// might be a bad idea in perspective since it does not return error
func WriteStruct(v any, w http.ResponseWriter, r *http.Request) {
	msg, _ := json.Marshal(v)
	w.Write(msg)
}