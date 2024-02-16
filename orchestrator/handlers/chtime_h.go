package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orchestrator/config"
)

func ChtimeHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var data ChtimeReqIn
	err := decoder.Decode(&data)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		e := ChtimeReqOut{"invalid request parameters"}
		WriteStruct(e, w, r)
		return
	}
	err = data.chsign()
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		e := ChtimeReqOut{err.Error()}
		WriteStruct(e, w, r)
		return
	}

	e := ChtimeReqOut{""}
	msg, _ := json.Marshal(e)
	w.Write(msg)
}

func (c *ChtimeReqIn) chsign() error{
	if c.Ms < 0{
		return fmt.Errorf("operation timeout cannot be less than 0")
	}
	sign := c.Sign
	if sign == "plus" || sign == "+"{
		config.Conf.Signs.Lock.Lock()
		config.Conf.Signs.Plus = c.Ms
		config.Conf.Signs.Lock.Unlock()
	} else if sign == "minus" || sign == "-" {
		config.Conf.Signs.Lock.Lock()
		config.Conf.Signs.Minus = c.Ms
		config.Conf.Signs.Lock.Unlock()
	} else if sign == "mul" || sign == "*" {
		config.Conf.Signs.Lock.Lock()
		config.Conf.Signs.Mul = c.Ms
		config.Conf.Signs.Lock.Unlock()
	} else if sign == "div" || sign == "/" {
		config.Conf.Signs.Lock.Lock()
		config.Conf.Signs.Div = c.Ms
		config.Conf.Signs.Lock.Unlock()
	}else {
		return fmt.Errorf("invalid sign %s", c.Sign)
	}
	return nil
}