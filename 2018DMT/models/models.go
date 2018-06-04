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
	Id   int    `gorm:"primary_key"`
	Name string `json:"Name"`
}
