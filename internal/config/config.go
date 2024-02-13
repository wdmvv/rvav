package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

var signs signConfig
var Conf Config

//Read config and parse it into Config structure (which must be accessed later, not reinited)
func NewConfig(path string) error{
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
	mp["+"] = signs.Plus
	mp["-"] = signs.Minus
	mp["*"] = signs.Mul
	mp["/"] = signs.Div

	Conf.Signs = mp
	return nil
}