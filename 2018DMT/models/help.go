package models

import "time"

type HelpReply struct {
	Id         int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	SeekHelpId int       `json:"SeekHelpId"`
	HelperId   int       `json:"HelperId"`
	HelperName string    `json:"HelperName"` //+
	HelperHead string    `json:"HelperHead"` //+
	Time       time.Time `json:"Time"`
	Context    string    `gorm:"type:varchar(1023)" json:"Context"`
	Accept     int       `json:"Accept"`
}
