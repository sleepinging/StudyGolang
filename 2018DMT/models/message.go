package models

type Message struct {
	Id      int    `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	Type    int    `json:"Type"`
	Title   string `json:"Title"`
	Context string `gorm:"type:varchar(2047)" json:"Context"`
	Readed  int    `json:"Readed"`
}
