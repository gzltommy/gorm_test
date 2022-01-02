package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test/tools"
)

type Player struct {
	gorm.Model
	ID       int64  // 主键
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

// TableName 设置表名，可以通过给 struct 类型定义 TableName 函数，返回当前 struct 绑定的 mysql 表名是什么
func (u Player) TableName() string {
	//绑定 MYSQL 表名为 users
	return "players"
}

func main() {
	//获取DB
	db := tools.GetDB()

	defer tools.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Player{})

	//执行数据库插入操作
	p := Player{
		Username: "tommy",
		Password: "123456",
	}

	if err := db.Create(&p).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}
