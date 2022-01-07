package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test/tools"
)

type User struct {
	ID         int64  // 主键
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func main() {
	// 获取DB
	db := tools.GetDB()
	defer tools.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	var user User
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	//stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	//stmt.Vars         //=> []interface{}{1}

	fmt.Println("===", stmt.SQL.String(), stmt.Vars)

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("id desc").Find(&[]User{})
	})
	fmt.Println("****", sql)
}
