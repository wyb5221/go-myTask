package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义结构体
type User struct {
	ID           uint           //主键的标准字段
	Name         string         //字符串，默认空字符串
	Email        *string        //指针，允许为空，默认null
	Age          uint8          //整数
	MemberNumber sql.NullString //允许为空，默认为null,可以通过字段Valid=true来设置保存空字符串
	ActivatedAt  sql.NullTime   //允许为空的时间
	Birthday     *time.Time
	CreateAt     time.Time `gorm:"autoCreateTime"` // 自动设置创建时间
	UpdateAt     time.Time `gorm:"autoUpdateTime"` // 自动更新时间
	Ignored      string
}

// Member表，使用gorm预定义的默认字段
type Member struct {
	gorm.Model
	Name   string
	Age    uint8
	UserNo string
}

type Author struct {
	Name  string
	Email string
}
type Blog struct {
	//嵌入结构体
	Author
	//嵌入结构体等价于引入字段
	// Name  string
	// Email string
	ID      int
	Upvotes int32
}

/**
 * 通过标签指定是嵌入Author
 */
// type Blog struct {
//   ID      int
//   Author  Author `gorm:"embedded"`
// 通过标签修改字段名称
//   Upvotes int32 `gorm:column:votes`
// }

/**
 * 创建数据库表
 */
func Run(db *gorm.DB) {
	// //创建数据库表
	// db.AutoMigrate(&User{})
	user := &User{}
	//插入一条空数据
	// db.Create(user)
	user.Name = "Go"
	email := "123321@qq.com"
	user.Email = &email
	user.Age = 10
	user.MemberNumber.Valid = true
	//插入一条数据
	db.Create(user)

	//创建数据库表
	// db.AutoMigrate(&Blog{})

}

func main() {
	fmt.Println("----")
	//创建数据库连接
	//设置数据库字符集：charset=utf8mb4
	//允许时间格式字段转成时间parseTime=True
	db, err :=
		gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	Run(db)
	fmt.Println("----")
}
