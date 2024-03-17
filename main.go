package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:8889)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.Create(&User{Username: "melih", Password: "selambenmelihinsifresi"})

	//var user User
	//db.First(&user, 1)
	//db.First(&user, "username = ?", "admin")
	//fmt.Println(user.Username, user.Password)

	//var users []User
	//db.Find(&users)
	//for _, user := range users {
	//	println(user.Username, user.Password)
	//}

	//var user User
	//db.First(&user, 1)
	////db.Model(&user).Update("Password", "newpassword")
	//db.Model(&user).Updates(User{Username: "newusername", Password: "newpassword"})

	var user User
	db.First(&user, 1)
	db.Delete(&user)
}
