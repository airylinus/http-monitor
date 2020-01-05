package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smf8/http-monitor/model"
	"strings"
)

func Setup(databaseName string) *gorm.DB {
	db := newDB(databaseName)
	migrate(db)
	// setup Database loading mode and connection pool and other settings as you prefer
	//db.DB().SetMaxIdleConns(5)
	return db
}
func newDB(name string) *gorm.DB {
	if !strings.HasSuffix(name, ".db") {
		name = name + ".db"
	}
	db, err := gorm.Open("sqlite3", "./"+name)
	if err != nil {
		fmt.Println("Error in creating database file : ", err)
		return nil
	}
	//Enable logging
	db.LogMode(true)
	return db
}
func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Request{}, &model.Url{})
}