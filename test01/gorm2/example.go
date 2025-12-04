package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	UserNo   string
	Age      uint
	Birthday time.Time
}

func Run(db *gorm.DB) {
	//创建表
	db.AutoMigrate(&User{})
	// user := User{Name: "张三", UserNo: "zhangsan", Age: 22, Birthday: time.Now()}
	// //插入数据需要使用指针
	// db.Create(&user)
	// //通过数据的指针数组来批量创建
	// users := []*User{
	// 	{Name: "李斯", UserNo: "lisi", Age: 18, Birthday: time.Now()},
	// 	{Name: "诸葛", UserNo: "zhuge", Age: 19, Birthday: time.Now()},
	// }
	// db.Create(users)
	// user := User{Name: "香香", UserNo: "xiangxiang", Age: 20, Birthday: time.Now()}
	// db.Select("Name", "Age", "CreatedAt").Create(&user)

	// user := User{}
	// // 获取第一条记录（主键升序）
	// db.First(&user)
	// //加debug可以打印执行的sql语句
	// db.Debug().First(&user)
	// fmt.Println("First--user1:", user)
	// // SELECT * FROM users ORDER BY id LIMIT 1;

	// // 获取一条记录，没有指定排序字段
	// db.Take(&user)
	// fmt.Println("Take--user2:", user)
	// // SELECT * FROM users LIMIT 1;
	// user := User{}
	// db.First(&user, 3)
	// fmt.Println("--user:", user)
	// user1 := User{}
	// user1.ID = 2
	// db.First(&user1)
	// fmt.Println("--user1:", user1)
	// db.First(&user, "id = ?", "2")
	// fmt.Println("--user:", user)
	// user2 := User{ID: 2}
	// db.First(&user2)

	// var user User
	// db.Debug().First(&user, 2)
	// db.Debug().First(&user, "2")
	// user.ID = 2
	// db.Debug().First(&user)
	// var user User
	// db.Debug().Find(&user, []int{1, 2, 3})
	// fmt.Println("--user:", user)

	var users []User
	db.Debug().First(&users)
	// fmt.Println("--user:", users)
}

func main() {
	fmt.Println("---")

	db, err :=
		gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	Run(db)
	fmt.Println("----")

}
