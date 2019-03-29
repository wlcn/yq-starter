package helper

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// just call it`s init method
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitAppDatabase 初始化数据库链接
func InitAppDatabase() error {
	var err error
	LogAccess.Infof("Init App Database Engine as %v", AppConfig.Database.Engine)
	switch AppConfig.Database.Engine {
	case "mysql":
		DB, err = gorm.Open("mysql", AppConfig.Database.Mysql.URL)
	case "sqlite3":
		DB, err = gorm.Open("sqlite3", AppConfig.Database.Sqlite3.URL)
	default:
		LogError.Error("Database error: can't find database driver")
		err = fmt.Errorf("can't find Database driver")
	}
	if err != nil {
		LogError.Errorf("Database error %+v", err)
		log.Fatalf("db err %+v", err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.LogMode(AppConfig.Database.LogMode)
	return err
}
