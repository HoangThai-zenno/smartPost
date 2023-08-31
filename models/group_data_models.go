package models

type Group struct {
	Id        int    `json:"id" gorm:"column:id" gorm:"primaryKey"`
	NameGroup string `json:"name_group" gorm:"column:name_group"`
	Active    bool   `json:"active" gorm:"column:active" gorm:"default:false"`
	Users     []User `gorm:"many2many:user_groups;"`
}
