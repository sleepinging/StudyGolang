package models

import "time"

type VisitRecord struct {
	Time time.Time
	IP   string
	Page string //访问的页面
}

type LoginRecord struct {
	UserId int
	Time   time.Time
	IP     string
	Su     int
}
