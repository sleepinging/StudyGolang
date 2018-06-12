package models

import (
	"time"
	"encoding/json"
)

type BlogReply struct {
	Id          int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	BlogId      int       `json:"BlogId"`
	ReplyerId   int       `json:"ReplyerId"`
	ReplyerName string    `json:"ReplyerName"` //+
	ReplyerHead string    `json:"ReplyerHead"` //+
	AtId        int       `json:"AtId"`
	AtName      string    `json:"AtName"` //+
	Time        time.Time `json:"Time"`
	Content     string    `gorm:"type:varchar(1023)" json:"Content"`
}

func LoadBlogReplyFromStr(str string) (reply *BlogReply, err error) {
	reply = new(BlogReply)
	err = json.Unmarshal([]byte(str), reply)
	return
}
