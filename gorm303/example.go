package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserLan struct {
	gorm.Model
	//多对多，需要通过标签指定关联关系many2many， 指定中间表的名称user_languages
	Languages []Language `gorm:"many2many:user_languages;"`
}
type Language struct {
	gorm.Model
	Name string
}

func Run(db *gorm.DB) {
	//生成表
	db.AutoMigrate(&UserLan{})
	db.AutoMigrate(&Language{})

	//插入用户
	// user := UserLan{}
	// db.Create(&user)

	// language1 := Language{Name: "english"}
	// db.Create(&language1)
	// //
	// language2 := Language{Name: "chinese"}
	// db.Create(&language2)
	//插入用户的时候会同步插入关联表Language的数据
	// user := UserLan{Languages: []Language{{Name: "en1"}, {Name: "en2"}}}
	// db.Create(&user)

	// var l []Language
	// db.Where("name in ?", []string{"en1", "en2"}).Find(&l)
	// 如果语言不存在，则创建
	// if len(l) == 0 {
	// 	l = []Language{
	// 		{Name: "en1"},
	// 		{Name: "en2"},
	// 	}
	// 	db.Create(&l)
	// }
	// // 使用查询到的语言记录
	// user2 := UserLan{Languages: l}
	// db.Create(&user2)
	// user3 := UserLan{Languages: l}
	// db.Create(&user3)

	user := UserLan{}
	err := db.Preload("Languages").Find(&user, 5).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

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
