package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserCred struct {
	gorm.Model
	//一对多
	CreditCards []CreditCard
}
type CreditCard struct {
	gorm.Model
	Number     string
	UserCredID uint
}

func Run(db *gorm.DB) {
	//生成表
	// db.AutoMigrate(&UserCred{})
	// db.AutoMigrate(&CreditCard{})

	//插入用户
	// user := UserCred{}
	// db.Create(&user)
	//插入卡信息
	// card := CreditCard{Number: "11111", UserCredID: 1}
	// db.Create(&card)

	//直接查询只有用户信息
	// u := UserCred{}
	// db.First(&u, 1)
	// fmt.Println("--u:", u)
	//使用Preload预加载，会查询出关联表的信息
	// db.Preload("CreditCard").First(&u, 1)
	// fmt.Println("--u:", u)

	//在插入卡
	// card2 := CreditCard{Number: "22222", UserCredID: 1}
	// db.Create(&card2)

	user := UserCred{}
	db.Preload("CreditCards").First(&user, 1)
	fmt.Println(user)
}

func main() {
	fmt.Println("---开始---")
	db, err :=
		gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	Run(db)
}
