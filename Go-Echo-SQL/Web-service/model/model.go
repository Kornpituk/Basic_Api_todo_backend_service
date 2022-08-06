package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Task  string `json:"Task"`
}
