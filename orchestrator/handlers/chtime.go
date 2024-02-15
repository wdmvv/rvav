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
		msg, _ := json.Marshal(e)
		w.Write(msg)
		return
	}
	err = data.chsign()
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		e := ChtimeReqOut{err.Error()}
		msg, _ := json.Marshal(e)
		w.Write(msg)
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
	switch c.Sign{
	case "plus":
		config.RWLock.RLock()
		config.Conf.Signs["plus"] = c.Ms
		config.RWLock.Unlock()
	case "minus":
		config.RWLock.RLock()
		config.Conf.Signs["minus"] = c.Ms
		config.RWLock.Unlock()
	case "mul":
		config.RWLock.RLock()
		config.Conf.Signs["mul"] = c.Ms
		config.RWLock.Unlock()
	case "div":
		config.RWLock.RLock()
		config.Conf.Signs["div"] = c.Ms
		config.RWLock.Unlock()
	default:
		return fmt.Errorf("invalid sign %s", c.Sign)
	}
	return nil
}