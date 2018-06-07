package models

import "time"

type BlogReply struct {
	Id          int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	BlogId      int       `json:"BlogId"`
	ReplyerId   int       `json:"ReplyerId"`
	ReplyerName string    `json:"ReplyerName"` //+
	ReplyerHead string    `json:"ReplyerHead"` //+
	AtId        int       `json:"AtId"`
	AtName      string    `json:"AtName"` //+
	Time        time.Time `json:"Time"`
	Context     string    `gorm:"type:varchar(1023)" json:"Context"`
}
