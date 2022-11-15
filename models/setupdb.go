package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// konfigurasi db mysql
	db, err := gorm.Open(mysql.Open("root:@tcp(mysql)/gateway_warehouse?parseTime=true"))
	//sqlDB, err := sql.Open("mysql", "root:''@tcp(mysql)/gateway_warehouse?allowOldPasswords=true")
	//db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})

	if err != nil {
		logrus.Println("Connection failed", err)
		panic(err)
	} else {
		logrus.Println("Connection established")
	}

	db.AutoMigrate(&TallySheet{}, &TallyTable{})
	DB = db
	logrus.Info("berhasil connect ke database")
}
