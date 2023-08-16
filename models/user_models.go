package models

import (
	"time"
)

type User struct {
	//gorm.Model
	Id         int        `json:"id" gorm:"column:id"`
	Name       string     `json:"name" gorm:"column:name" `
	Email      string     `json:"email"gorm:"column:email"`
	Time_Stamp *time.Time `json:"time_Stamp" gorm:"column:time_stamp"`
	Role       string     `json:"role" gorm:"column:role"`
	Avatar     string     `json:"avatar" gorm:"column:avatar"`
	Groups     string     `json:"groups" gorm:"column:groups"`
}
