package model

import "github.com/jinzhu/gorm"

type DatabaseConf struct {
	DBType   string `json:"db_type"`
	Host     string `json:"db_host"`
	Port     string `json:"db_port"`
	DBname   string `json:"db_dbname"`
	Username string `json:"db_username"`
	Password string `json:"db_password"`
	Charset  string `json:"db_charset"`
}

type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(20);not null"`
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(32);not null"`
	YepID    string `gorm:"type:int(9);not null"`
}
