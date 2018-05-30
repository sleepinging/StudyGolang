package models

type RetJson struct {
	Status int    `json:status`
	Msg    string `json:msg`
	Data   string `json:data`
}
