package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	UserNo   string
	Age      uint8
	Job      string
	Birthday sql.NullTime
}

func Run(db *gorm.DB) {
	// var user User
	//创建数据库表
	// db.AutoMigrate(&user)
	//单条插入
	// u1 := User{Name: "李一", UserNo: "liyi01", Age: 20, Job: "金牌销售", Birthday: sql.NullTime{
	// 	Time:  time.Date(2004, 5, 15, 0, 0, 0, 0, time.UTC), // 2004年5月15日
	// 	Valid: true,
	// }}
	// db.Create(&u1)
	// u2 := User{Name: "李靖", UserNo: "lijin", Age: 35, Job: "高级战神", Birthday: sql.NullTime{
	// 	Time:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), // 2004年5月15日
	// 	Valid: true,
	// }}
	// r := db.Create(&u2)
	// fmt.Println("--r:", r)
	// a := r.RowsAffected
	// fmt.Println("--影响行数a:", a)
	// fmt.Println("--返回的实体u2:", u2)
	// fmt.Println("--返回的实体id:", u2.ID)

	// var users = []User{{Name: "大乔", UserNo: "daqiao", Age: 20, Job: "高级导师", Birthday: sql.NullTime{
	// 	Time:  time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Valid: true,
	// }}, {Name: "小乔", UserNo: "xiaoqiao", Age: 18, Job: "导师", Birthday: sql.NullTime{
	// 	Time:  time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Valid: true,
	// }}, {Name: "香香", UserNo: "xiangxiang", Age: 35, Job: "测试", Birthday: sql.NullTime{
	// 	Time:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Valid: true,
	// }}, {Name: "飞飞", UserNo: "feifei", Age: 35, Job: "财务", Birthday: sql.NullTime{
	// 	Time:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Valid: true,
	// }}}
	// r := db.Create(&users)
	// fmt.Println("--r:", r)

	// u1 := User{UserNo: "xiaoqiao"}
	//查询所有数据
	// var s []User
	// db.Debug().Find(&s)
	// fmt.Println("--s:", s)
	// db.Debug().Where("name=?", "香香").Find(&s)
	// fmt.Println("--s:", s)
	// db.Debug().Where("user_no=?", "zhuge").Find(&s)
	// fmt.Println("--s:", s)
	// db.Debug().Where("name in ?", []string{"李靖", "小乔"}).Find(&s)
	// fmt.Println("--s:", s)
	// db.Debug().Where("name <> ?", "李二").Find(&s)
	// fmt.Println("--s:", s)
	// db.Debug().Where("name like ?", "%李%").Find(&s)
	// fmt.Println("--s:", s)
	// db.Debug().Where("user_no like ? and age>?", "%i%", 20).Find(&s)
	// fmt.Println("--s:", s)
	var s User
	db.Model(&s).Where("name", "李二").Update("age", "22")
	db.Model(&s).Where("id=?", "3").Update("age", "22")
	s.ID = 6
	db.Model(&s).Update("age", "22")
	db.Debug().Model(&s).Where("user_no", "xiaoqiao").Update("age", "22")
	// 使用结构体更新多列
	db.Model(&s).Updates(User{Name: "大乔1", UserNo: "daqiao1", Age: 19})
	// 使用map更新多列
	db.Model(&s).Updates(map[string]interface{}{"name": "小乔1", "user_no": "xiaoqiao1", "age": 18})
	// 指定更新name字段
	db.Debug().Model(&s).Select("name").Updates(map[string]interface{}{"name": "小乔2", "user_no": "xiaoqiao2", "age": 11})
	// 指定更新name字段
	db.Debug().Model(&s).Select("name", "age").Updates(map[string]interface{}{"name": "小乔2", "user_no": "xiaoqiao2", "age": 0})
	db.Debug().Model(&s).Select("name", "age").Updates(User{Name: "小乔3", UserNo: "xiaoqiao3", Age: 0})
}

func main() {
	fmt.Println("--开始--")
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	Run(db)
}
