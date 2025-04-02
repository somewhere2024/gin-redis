package models

// 数据库models
type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;unique;not null;autoIncrement"`
	UUId     string `json:"uuid" gorm:"size:100;not null"`
	Username string `json:"username" gorm:"size:100;not null"`
	Password string `json:"password" gorm:"size:100;not null"`
}

type Sub struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type UserLoginIn struct {
	username string
	password string
	UUId     string
}
