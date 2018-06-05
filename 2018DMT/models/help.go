package models

import "time"

//帮助
type Help struct {
	Id      int       `gorm:"primary_key" json:"Id"`
	UserId  int       `json:"UserId"`
	Type    int       `json:"Type"`
	Title   string    `json:"Title"`
	Context string    `gorm:"type:varchar(4095)" json:"Context"`
	Gold    int       `json:"Gold"`
	Time    time.Time `json:"Time"`
	Status  int       `json:"Status"`
}
