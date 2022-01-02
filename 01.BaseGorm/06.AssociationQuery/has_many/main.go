package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test/tools"
)

type User struct {
	gorm.Model
	// 通过标签，将外键定义为：UserRefer
	CreditCards []CreditCard `gorm:"foreignkey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint // 新定义的外键名
}

func main() {
	//获取DB
	db := tools.GetDB()
	defer tools.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&CreditCard{})

	user := User{}
	// 查询用户数据
	//自动生成 sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	db.Where("username = ?", "tizi365").First(&user)
	fmt.Println(user)

	//自动生成SQL： SELECT * FROM emails WHERE user_id = 111; // 111 是 user 的主键 ID 值
	// 关联查询的结果，保存到 user.CreditCard 属性
	db.Model(&user).Association("CreditCard").Find(&user.CreditCards)
}
