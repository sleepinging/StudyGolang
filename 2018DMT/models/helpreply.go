package models

import "time"

type HelpReply struct {
	Id      int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	BlogId  int       `json:"BlogId"`
	UserId  int       `json:"UserId"`
	AtId    int       `json:"AtId"`
	Time    time.Time `json:"Time"`
	Context string    `gorm:"type:varchar(1023)" json:"Context"`
	Accept  int       `json:"Accept"`
}
