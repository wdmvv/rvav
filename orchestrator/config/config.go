package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type signConfig struct{
	Plus int `json:"plus"`
	Minus int `json:"minus"`
	Mul int `json:"mul"`
	Div int `json:"div"`
}

type Config struct{
	Signs map[string]int `json:"-"`
	Workers int `json:"workers"`
	Timeout int `json:"timeout"`
}

var RWLock sync.RWMutex // need it to change timeouts with no unexpected outcomes
var signs signConfig
var Conf Config

//Read config and parse it into Config structure (which must be accessed later, not reinited)
func NewConfig(path string) error{
	RWLock = sync.RWMutex{}
	wd, err := os.Getwd()
	if err != nil{
		return err
	}
	path = filepath.Join(wd, path)
	f, err := os.Open(path)
  	if err != nil{
	    return fmt.Errorf("failed to read config file %s", path)
  	}
  	data, err := io.ReadAll(f)
  	if err != nil{
    	return fmt.Errorf("failed to read config file %s", path)
	}
	err = json.Unmarshal(data, &signs)
	if err != nil{
		return fmt.Errorf("failed to parse config %s", path)
	}
	err = json.Unmarshal(data, &Conf)
	if err != nil{
		return fmt.Errorf("failed to parse config %s", path)
	}

	mp := make(map[string]int)
	mp["plus"] = signs.Plus
	mp["minus"] = signs.Minus
	mp["mul"] = signs.Mul
	mp["div"] = signs.Div

	Conf.Signs = mp
	return nil
}
