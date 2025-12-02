package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       uint8
	CompanyId int
	Company   Company
}

type Company struct {
	gorm.Model
	Name       string
	CardNumber int `gorm:"unique"`
}

func Run(db *gorm.DB) {
	//生成表
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Company{})

	//插入数据
	// c := Company{Name: "北京科技有限公司", CardNumber: 1001}
	// db.Create(&c)
	// u := User{Name: "张三", Age: 25, CompanyId: 1}
	// db.Create(&u)

	//查询，只会返回user对象信息
	// var u User
	// db.First(&u)
	// fmt.Println("查询返回：", u)
	//Preload预加载，会返回关联表对象信息
	var u User
	db.Preload("Company").First(&u)
	fmt.Println("查询返回：", u)
	fmt.Println("查询返回：", u.Name, u.Company.Name)
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
