package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	First_name string `json:"First_name"`
	Last_name string `json:"Last_name"`
	Task  string `json:"Task"`
}
