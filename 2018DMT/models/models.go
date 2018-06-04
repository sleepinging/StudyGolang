package models

type RetJson struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type Login struct {
	Email    string `gorm:"primary_key" json:"email"`
	Password string `json:"password"`
}

type Job struct {
	Id          int     `gorm:"primary_key"`
	Name        string  `json:"Name"`
	Salary      float32 `json:"Salary"`
	Time        string  `json:"Time"`
	Weekend     int     `json:"Weekend"`
	Pickup      int     `json:"Pickup"`
	Eat         int     `json:"Eat"`
	Live        int     `json:"Live"`
	WuXianYiJin int     `json:"WuXianYiJin"`
	Place       string  `json:"Place"`
	LimPeople   int     `json:"LimPeople"`
	NowPeople   int     `json:"NowPeople"`
	Sex         int     `json:"Sex"`
	Phone       string  `json:"Phone"`
	Detail      string  `json:"Detail"`
}
