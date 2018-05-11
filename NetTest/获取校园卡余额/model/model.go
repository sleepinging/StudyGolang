package model

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Openid string `json:"openid"`
	SMTP   Smtp   `json:"smtp"`
	Limit  int    `json:"limit"`
}

type Smtp struct {
	Pwd    string `json:"pwd"`
	Host   string `json:"host"`
	To     string `json:"to"`
	Sender string `json:"sender"`
}

func (cfg *Config) Load(filename string) (err error) {
	//fmt.Println(filename)
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(dat, cfg)
	//fmt.Println(string(dat))
	return
}
