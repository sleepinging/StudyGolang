package models

import (
	"encoding/json"
	"../tools"
)

type User struct {
	Id    int    `gorm:"primary_key" json:"Id"`
	Email string `json:"email"`
}

func (this *User) ToString() (str string) {
	b, err := json.Marshal(this)
	tools.ShowErr(err)
	str = string(b)
	return
}
