package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Balance uint64
}

type Transactions struct {
	gorm.Model
	From_account_id uint
	To_account_id   uint
	Amount          uint64
}

// 创建数据库表
func CreateTable02(db *gorm.DB) {
	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transactions{})
}

// 通过数组批量插入数据
func BtaInsert02(db *gorm.DB) {
	var s = []Accounts{{
		Balance: 1000,
	}, {
		Balance: 10,
	}, {
		Balance: 2000,
	}}
	r := db.Create(&s)
	fmt.Println("成功插入条数：", r.RowsAffected)
	fmt.Println("成功插入的数据集：", s)
}

// 转账
func Transfer(db *gorm.DB, num uint, fId uint, tId uint) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var account Accounts
		account.ID = fId
		tx.Debug().First(&account)
		fmt.Println("--account:", account)
		if account.Balance < uint64(num) {
			return errors.New("账户余额不足")
		}

		err2 := tx.Debug().Model(&Accounts{}).Where("id=?", tId).Update("balance", gorm.Expr("balance+?", num)).Error
		if err2 != nil {
			return err2
		}

		err1 := tx.Debug().Model(&account).Update("balance", account.Balance-uint64(num)).Error
		if err1 != nil {
			return err1
		}
		t := Transactions{
			From_account_id: fId,
			To_account_id:   tId,
			Amount:          uint64(num),
		}
		err3 := tx.Create(&t).Error
		if err3 != nil {
			return err3
		}
		// return nil
		return errors.New("【模拟异常】系统故障，强制回滚")
	})
	fmt.Println("-err:", err)
	if err != nil {
		fmt.Println("转账失败")
	} else {
		fmt.Println("转账成功")
	}
}

// func main() {
// 	fmt.Println("---开始---")
// 	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	// CreateTable02(db)
// 	// BtaInsert02(db)
// 	Transfer(db, 200, 3, 2)

// }
