package helper

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// just call it`s init method
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	Once.Do(Init)
}

// Init 初始化数据库链接
func Init() {
	tmpDir := "/tmp/yq-starter/"
	os.MkdirAll(tmpDir, os.ModePerm)
	// os.RemoveAll(tmpDir + "app_test.db")
	db, err := gorm.Open("sqlite3", tmpDir+"app_test.db")
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	db.LogMode(true)
	DB = db
}

// Close 关闭数据库
func Close() {
	DB.Close()
}
