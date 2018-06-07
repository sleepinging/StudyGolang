package models

import "time"

type Message struct {
	Id         int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	SenderId   int       `json:"SenderId"`
	SenderName string    `json:"SenderName"` //+
	RecverId   int       `json:"RecverId"`
	RecverName string    `json:"RecverName"` //+
	Time       time.Time `json:"Time"`
	Type       int       `json:"Type"`
	Title      string    `json:"Title"`
	Context    string    `gorm:"type:varchar(2047)" json:"Context"`
	Readed     int       `json:"Readed"`
}
