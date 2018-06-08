package models

import "time"

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
