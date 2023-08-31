package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id" gorm:"column:id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"column:name" `
	Email     string     `json:"email" gorm:"column:email"`
	TimeStamp *time.Time `json:"time_Stamp" gorm:"column:time_stamp"`
	Role      string     `json:"role" gorm:"column:role"`
	Avatar    string     `json:"avatar" gorm:"column:avatar"`
	Groups    []Group    `gorm:"many2many:user_groups;"`
}
