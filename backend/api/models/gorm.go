package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
func GormDB() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME   := "sample"
	TIME 	 := "?parseTime=true"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME + TIME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&User{})
	return db
}