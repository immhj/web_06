package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"web_06/models"
)

var DB *gorm.DB

func Initmysql() (err error) {
	DB, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db_06?charset=utf8mb4&parseTime=True&loc=Local")
	DB.AutoMigrate(&models.User{})
	return err
}

func Selectuser(username string) (ok bool) {
	user := models.User{}
	DB.AutoMigrate(&models.User{})
	DB.Where("username = ?", username).First(&user)
	if len(user.Username) == 0 {
		return true
	}
	return false
}

func Selectpasswordbyusername(username string) (password string) {
	user := models.User{}
	DB.Where("username = ?", username).First(&user)
	return user.Password
}

func Createuser(username, password string) {
	user := models.User{}
	user.Username = username
	user.Password = password
	DB.Create(&user)
}
