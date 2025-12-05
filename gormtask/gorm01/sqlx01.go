package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 结构体字段使用指针允许字段为空
type EmployeeInfo struct {
	Id         uint
	Name       string
	Department string
	Salary     uint
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *string    `db:"deleted_at"`
}

type Books struct {
	Id     uint
	Title  string
	Author string
	Price  string
}

func insertEm1(db *sqlx.DB, employee EmployeeInfo) {
	sql1 := "INSERT INTO employees (name,department,salary) VALUES (?, ?, ?)"
	db.Exec(sql1, employee.Name, employee.Department, employee.Salary)
}
func insertEm2(db *sqlx.DB, employee EmployeeInfo) {
	sql2 := "INSERT INTO employees (name,department,salary) VALUES (:name, :department, :salary)"
	result, err := db.NamedExec(sql2, employee)
	if err != nil {
		panic(err)
	}
	num, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	fmt.Printf("插入成功:%d，ID: %d\n", num, id)
}
func insertEm2Optimized(db *sqlx.DB, employee EmployeeInfo) error {
	// 使用 PrepareNamed 预编译 SQL，提高性能
	stmt, err := db.PrepareNamed(`
		INSERT INTO employees (name, department, salary) 
		VALUES (:name, :department, :salary)
	`)
	if err != nil {
		return fmt.Errorf("准备语句失败: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(employee)
	if err != nil {
		return fmt.Errorf("执行插入失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取插入ID失败: %w", err)
	}

	fmt.Printf("插入成功，ID: %d\n", id)
	return nil
}

func batchInsertEmployees(db *sqlx.DB, employees []EmployeeInfo) error {
	sql := "INSERT INTO employees (name, department, salary) VALUES (:name, :department, :salary)"

	// 使用 NamedExec 批量插入
	result, err := db.NamedExec(sql, employees)
	if err != nil {
		return fmt.Errorf("批量插入失败: %w", err)
	}
	num, _ := result.RowsAffected()
	fmt.Printf("批量插入成功，插入了 %d 条记录\n", num)
	return nil
}

// 查询
func GetEmployeesById(db *sqlx.DB, id uint) *EmployeeInfo {
	var e EmployeeInfo
	sql := "select * from employees where id=?"
	err := db.Get(&e, sql, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return &e
}

func GetEmployeesByDepartment(db *sqlx.DB, department string) *[]EmployeeInfo {
	var e []EmployeeInfo
	sql := "select * from employees where department = ?"
	err := db.Select(&e, sql, department)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return &e
}

func GetEmployeesByDepartments(db *sqlx.DB, department string) *[]EmployeeInfo {
	var e []EmployeeInfo
	sql := "select * from employees where department like ?"
	err := db.Select(&e, sql, "%"+department+"%")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return &e
}

func GetEmployeesBySalarys(db *sqlx.DB) *[]EmployeeInfo {
	var e []EmployeeInfo
	sql := "select * from employees where salary = (select MAX(salary) from employees) order by id "
	err := db.Select(&e, sql)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return &e
}

func updateSalaryById(db *sqlx.DB, salary uint, id uint) {
	sql := "update employees set salary=? where id=?"
	result, err := db.Exec(sql, salary, id)
	if err != nil {
		panic(err)
	}
	n, _ := result.RowsAffected()
	fmt.Println("成功执行条数：", n)
}

func queryBookByPrice(db *sqlx.DB, price uint) *[]Books {
	sql := "select * from books where price > ?"
	var b []Books
	err := db.Select(&b, sql, price)
	if err != nil {
		fmt.Printf("queryBookByPrice failed, err:%v\n", err)
		return nil
	}
	return &b
}

func initSqlxDb() sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 设置连接池参数
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	return *db
}

func initGormDb1() gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return *db
}

func main() {
	fmt.Println("---开始---")
	// db := initGormDb1()
	// db.AutoMigrate(&Books{})
	db := initSqlxDb()
	// fmt.Println("--GetEmployeesByDepartment:", GetEmployeesById(&db, 5))
	// fmt.Println("--GetEmployeesByDepartment:", GetEmployeesByDepartments(&db, "部"))
	// updateSalaryById(&db, 3300, 1)
	// GetEmployeesBySalarys(&db)
	// fmt.Println("--GetEmployeesByDepartment:", GetEmployeesBySalarys(&db))

	fmt.Println("--queryBookByPrice:", queryBookByPrice(&db, 70))
}
