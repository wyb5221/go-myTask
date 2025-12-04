package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Address struct {
	ID       uint `gorm:"primarykey"`
	Address1 string
	UserID   uint
}

type Email struct {
	ID     uint `gorm:"primarykey"`
	Email  string
	UserID uint
}

type Language struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type Company struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type User struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	BillingAddress Address //一对一（不能有两个关联关系的字段）
	// ShippingAddress Address
	Emails    []Email    //一对多
	Languages []Language `gorm:"many2many:user_languages;"` //多对多
	CompanyID uint       //多对一（需要现有主表数据）
	Company   Company
}

func CreateUser(db *gorm.DB) {
	comp := Company{Name: "北京科技有限公司"}
	db.Create(&comp)

	user := User{
		Name:           "张飞01",
		BillingAddress: Address{Address1: "Billing Address - Address 11111"},
		// ShippingAddress: Address{Address1: "Shipping Address - Address 111111"},
		Emails: []Email{
			{Email: "11@example.com"},
			{Email: "22@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
		CompanyID: comp.ID,
	}
	db.Create(&user)
}

func Run(db *gorm.DB) {
	//创建表
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Language{})
	// db.AutoMigrate(&Email{})
	// db.AutoMigrate(&Address{})
	// db.AutoMigrate(&Company{})

	// CreateUser(db)

	// var user User
	// db.First(&user)
	// fmt.Println("--user1:", user)
	// // Preload预加载指定的属性
	// db.Preload("BillingAddress").Preload("Emails").First(&user)
	// fmt.Println("--user2:", user)
	// // clause.Associations 预加载所有关联属性
	// db.Preload(clause.Associations).First(&user)
	// fmt.Println("--user3:", user)

	//只查询指定的关联属性
	// var langs []Language
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Find(&langs)
	// fmt.Println("--1:", langs)
	// db.Debug().Where(&User{ID: 1}).Association("Languages").Find(&langs)
	// fmt.Println("--2:", langs)
	// db.Debug().Where("id=?", 1).Association("Languages").Find(&langs)
	// fmt.Println("--3:", langs)
	//等价于
	// var user User
	// db.Preload("Languages").First(&user)
	// fmt.Println("--5:", user.Languages)

	// var user User
	// //先查询user属性
	// db.Preload(clause.Associations).First(&user, 1)
	// // user.BillingAddress = Address{Address1: "111"}
	// //在修改user属性的值
	// user.BillingAddress.Address1 = "99999"
	// //Updates不会更新关联表的非外键字段
	// // db.Debug().Updates(&user)
	// db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	// 如果user中有2个Address，就会发生异常，因为2个地址不能明确的判定出对应关系，这个时候应该在结构中明确指定地址ID

	// 一对多更新
	// var emails []Email
	// db.Model(&User{ID: 1}).Association("Emails").Find(&emails)
	// emails[0].Email = "1111@example.com"
	// db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Model(&User{ID: 1}).Association("Emails").Replace(emails)
	// db.Debug().Model(&User{ID: 1}).Association("Emails").Replace(&Email{Email: "11@11.com"}, &Email{Email: "22@22.com"})

	// 多对多更新
	var langZH, langEN Language
	db.First(&langZH, "name = ?", "ZH")
	db.First(&langEN, "name = ?", "EN")
	fmt.Println("langZH:", langZH)
	fmt.Println("langEN:", langEN)
	//Replace替换
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Replace(&langZH) // 必须是引用
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Delete(langZH)
	//Append 添加
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Append(&langEN) // 必须是引用
	//Append 添加一个没有的数据
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Append(&Language{Name: "FR"})
	//Clear清空
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Clear()

	// 删除关联
	db.Debug().Select("Emails", "Languages", "BillingAddress").Delete(&User{ID: 1})
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
