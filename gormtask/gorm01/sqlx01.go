package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type EmployeeInfo struct {
	Id         uint
	Name       string
	Department string
	Salary     uint
}

func insertEm1(db *sqlx.DB, employee EmployeeInfo) {
	sql1 := "INSERT INTO employees (name,department,salary) VALUES (?, ?, ?)"
	db.Exec(sql1, employee.Name, employee.Department, employee.Salary)
}
func insertEm2(db *sqlx.DB, employee EmployeeInfo) {
	sql2 := "INSERT INTO employees (name,department,salary) VALUES (:Name, :Department, :Salary)"
	db.Exec(sql2, employee)
}

// func GetEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {

// }

func initSqlxDb() sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return *db
}

func main() {
	fmt.Println("---开始---")
	// db := initGormDb()
	// queryEmployeeByDepart(&db, "技术部")
	db := initSqlxDb()
	// e1 := EmployeeInfo{
	// 	Name:       "李靖",
	// 	Department: "架构部",
	// 	Salary:     6000,
	// }
	e2 := EmployeeInfo{
		Name:       "李牧",
		Department: "架构部",
		Salary:     5500,
	}
	// insertEm1(&db, e1)
	insertEm2(&db, e2)
}
