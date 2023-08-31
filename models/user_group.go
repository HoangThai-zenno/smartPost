package models

type UserGroup struct {
	UserId  int `json:"UserId" gorm:"column:user_id"`
	GroupId int `json:"GroupId" gorm:"column:group_id"`
}
