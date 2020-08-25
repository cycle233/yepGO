package gdo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"strings"
	"yeeep.cn/yepgo/conf"
)

func InitDB() *gorm.DB {
	dbConf := conf.LoadDBConf()
	args := ""
	switch strings.ToLower(dbConf.DBType) {
	case "mysql":
		args = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			dbConf.Username,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.DBname,
			dbConf.Charset)
	case "postgres":
		args = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			dbConf.Host,
			dbConf.Port,
			dbConf.Username,
			dbConf.DBname,
			dbConf.Password)
	case "mssql":
		args = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			dbConf.Username,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.DBname)
	//case "sqlite3":
	default:
		panic("failed to recognized this database type")
	}
	db, err := gorm.Open(strings.ToLower(dbConf.DBType), args)

	if err != nil {
		logrus.Error(err.Error())
		panic("failed to connect database, err:" + err.Error())
	}
	logrus.Debug(db)
	return db
}
