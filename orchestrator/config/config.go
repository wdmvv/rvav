package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"strings"
)

type SignConfig struct{
	Plus int `json:"plus"`
	Minus int `json:"minus"`
	Mul int `json:"mul"`
	Div int `json:"div"`
	Lock sync.Mutex `json:"-"`
}

type Config struct{
	Signs *SignConfig
	Workers int `json:"workers"`
	Timeout int `json:"timeout"`
	AgentPort int `json:"agent_port"`
	//db config
	UseDB bool `json:"usedb"`
	User string `json:"user"`
	Pswd string `json:"password"`
	DBname string `json:"dbname"`
	TabName string `json:"tablename"`
}

var signs *SignConfig
var Conf *Config

//Read config and parse it into Config structure (which must be accessed later, not reinited)
func NewConfig(path string) error{
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
	signs.Lock = sync.Mutex{}
	err = json.Unmarshal(data, &Conf)
	if err != nil{
		return fmt.Errorf("failed to parse config %s", path)
	}
	Conf.Signs = signs
	
	return nil
}

// get sign timeout if it is valid
func (c *Config) SignTimeout(sign string) (int, error){
	sign = strings.ToLower(sign)
	var err error
	out := 0
	
	if sign == "plus" || sign == "+"{
		out = c.Signs.Plus
	} else if sign == "minus" || sign == "-"{
		out = c.Signs.Minus
	} else if sign == "mul" || sign == "*"{
		out = c.Signs.Mul
	} else if sign == "div" || sign == "/"{
		out = c.Signs.Div
	} else {
		err = fmt.Errorf("invalid sign %s", sign)
	}
	return out, err
}