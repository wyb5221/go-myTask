package main

import (
	_ "gorm.io/gorm"
)

// func main() {
// 	fmt.Println("---开始---")
// 	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	// db.AutoMigrate(&Employee{})
// 	var s = []Employee{{
// 		Name:       "香香",
// 		Department: "测试",
// 		Salary:     2000,
// 	}, {
// 		Name:       "张飞",
// 		Department: "技术部",
// 		Salary:     5000,
// 	}, {
// 		Name:       "赵云",
// 		Department: "技术部",
// 		Salary:     5500,
// 	}, {
// 		Name:       "刘备",
// 		Department: "技术部",
// 		Salary:     6000,
// 	}, {
// 		Name:       "大乔",
// 		Department: "财务部",
// 		Salary:     3000,
// 	}}
// 	db.Create(&s)

// 	// CreateTable02(db)5
// 	// BtaInsert02(db)
// 	// Transfer(db, 200, 3, 2)

// }
