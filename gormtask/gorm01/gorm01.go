package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	Name  string
	Age   uint8
	Grade string
}

// 创建数据库表
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Students{})
}

// 插入一条数据
func Insert(db *gorm.DB) {
	s := Students{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	r := db.Create(&s)
	fmt.Println("成功插入条数：", r.RowsAffected)
	fmt.Println("成功插入的数据集：", s)
}

// 通过数组批量插入数据
func BtaInsert(db *gorm.DB) {
	var s = []Students{{
		Name:  "李四",
		Age:   22,
		Grade: "四年级",
	}, {
		Name:  "王五",
		Age:   18,
		Grade: "二年级",
	}, {
		Name:  "刘六",
		Age:   21,
		Grade: "三年级",
	}}
	r := db.Create(&s)
	fmt.Println("成功插入条数：", r.RowsAffected)
	fmt.Println("成功插入的数据集：", s)
}

// 根据年龄查询大于入参的数据
func QueryByAge(db *gorm.DB, age int) []Students {
	var s []Students
	// db.Debug().Where("age>?", age).Find(&s)
	//结构体组装查询条件
	// db.Debug().Where(&Students{Age: uint8(age)}).Find(&s)
	//map组装查询条件
	db.Debug().Where(map[string]interface{}{"age": age}).Find(&s)
	fmt.Println("查询返回结果：", s)
	return s
}

// 修改
func UpdateByName(db *gorm.DB, name string, grade string) {
	var s Students
	// db.Debug().Model(&s).Where("name=?", name).Update("grade", grade)
	//结构体组装
	// db.Debug().Model(&s).Where(&Students{Name: name}).Updates(Students{Grade: grade})
	//map组装查询
	db.Debug().Model(&s).Where(map[string]interface{}{"name": name}).Updates(map[string]interface{}{"grade": grade})
	fmt.Println("修改后的信息：", s)
}

// 删除
func delByAge(db *gorm.DB, age int) {
	var s Students
	db.Debug().Where("age>?", age).Delete(&s)
}

func main() {
	fmt.Println("---开始---")
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	// CreateTable(db)
	// Insert(db)
	// BtaInsert(db)
	// QueryByAge(db, 21)
	// UpdateByName(db, "张三", "四年级2")
	delByAge(db, 21)
}
