package models

import "time"

type HelpReply struct {
	Id      int       `gorm:"primary_key" json:"Id"`
	BlogId  int       `json:"BlogId"`
	UserId  int       `json:"UserId"`
	AtId    int       `json:"AtId"`
	Time    time.Time `json:"Time"`
	Context string    `gorm:"type:varchar(1023)" json:"Context"`
	Accept  int       `json:"Accept"`
}
