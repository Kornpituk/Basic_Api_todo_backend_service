package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	//   Id  uint
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Task  string `json:"Task"`
	//   CreatedData string
	//   UpdatedData string
}
