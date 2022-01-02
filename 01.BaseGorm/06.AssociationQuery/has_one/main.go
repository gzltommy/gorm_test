package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test/tools"
)

//// 信用卡
//type CreditCard struct {
//	// 继承gorm的基础Model,里面默认定义了ID、CreatedAt、UpdatedAt、DeletedAt 4个字段
//	gorm.Model
//	Number   string
//	UserID   uint // 外键
//}
//
//// 用户
//type User struct {
//	gorm.Model
//	CreditCard   CreditCard // 持有信用卡属性（关联信用卡）
//}

type User struct {
	gorm.Model
	Name       string     `gorm:"index"`
	CreditCard CreditCard `gorm:"foreignkey:UserName;references:name"`
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

func main() {
	//获取DB
	db := tools.GetDB()
	defer tools.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	fmt.Println("-------------", db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&CreditCard{}).Error())

	//user := User{}
	//// 查询用户数据
	////自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	//db.Where("username = ?", "tizi365").First(&user)
	//fmt.Println(user)
	//
	//var card CreditCard
	//////自动生成 SQL： SELECT * FROM credit_cards WHERE user_id = 123; // 123 自动从user的ID读取
	//// 关联查询的结果会填充到 card 变量
	//db.Model(&user).Association("CreditCard").Find(&card)
}
