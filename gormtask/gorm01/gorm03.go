package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	Id         uint
	Name       string
	Department string
	Salary     uint
}

func queryEmployeeByDepart(db *gorm.DB, depart string) []Employee {
	var ess []Employee
	db.Debug().Where(&Employee{Department: depart}).Find(&ess)
	fmt.Println(ess)
	return ess
}

func initGormDb() gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return *db
}

// func main() {
// 	fmt.Println("---开始---")
// 	// db := initGormDb()
// 	// queryEmployeeByDepart(&db, "技术部")
// 	db := initSqlxDb()

// }
