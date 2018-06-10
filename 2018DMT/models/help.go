package models

import (
	"encoding/json"
	"time"
)

type Help struct {
	Id         int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	SeekHelpId int       `json:"SeekHelpId"`
	HelperId   int       `json:"HelperId"`
	HelperName string    `json:"HelperName"` //+
	HelperHead string    `json:"HelperHead"` //+
	Time       time.Time `json:"Time"`
	Content    string    `gorm:"type:varchar(1023)" json:"Content"`
	Accept     int       `json:"Accept"`
}

//从字符串加载
func LoadHelpFromStr(str string) (help *Help, err error) {
	help = new(Help)
	err = json.Unmarshal([]byte(str), help)
	return
}

type HelpReply struct {
	Id          int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	HelpId      int       `json:"HelpId"`
	ReplyerId   int       `json:"ReplyerId"`
	ReplyerName string    `json:"ReplyerName"`
	AtId        int       `json:"AtId"`
	AtName      string    `json:"AtName"`
	Time        time.Time `json:"Time"`
	Content     string    `gorm:"type:varchar(1023)" json:"Content"`
}

//从字符串加载
func LoadHelpReplyFromStr(str string) (hr *HelpReply, err error) {
	hr = new(HelpReply)
	err = json.Unmarshal([]byte(str), hr)
	return
}
