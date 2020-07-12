package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var GloalDb *gorm.DB

func InitDb() {
	GloalDb = NewDB()
}

// NewDB ...
func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", SConfig.Mysql)
	db.LogMode(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
