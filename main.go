package main

import (
	"fmt"
	example "go-myTask/gorm1"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// _ "go-myTask/task1"
	// _ "go-myTask/task2"
	// _ "go-myTask/task3"
	// task4 "go-myTask/task4"
)

func main() {
	fmt.Println("----")
	//创建数据库连接
	db, err :=
		gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	example.Run(db)
	fmt.Println("----")
}
