package main

import (
	"gorm.io/gorm"
	"gorm_test/tools"
)

// 用户表
type User struct {
	gorm.Model
	Username string
	Orders   []Orders // 关联订单，一对多关联关系
}

// 订单表
type Orders struct {
	gorm.Model
	UserID uint // 外键字段
	Price  float64
}

func main() {
	//获取DB
	db := tools.GetDB()
	defer tools.Close()

	var users []User
	// 预加载Orders字段值，Orders字段是User的关联字段
	db.Preload("Orders").Find(&users)
	// 下面是自动生成的SQL，自动完成关联查询
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4);

	// Preload第2，3个参数支持设置SQL语句条件和绑定参数
	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	// 自动生成的SQL如下
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4) AND state NOT IN ('cancelled');

	// 通过组合Where函数一起设置SQL条件
	db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	// 自动生成的SQL如下
	//// SELECT * FROM users WHERE state = 'active';
	//// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('cancelled');

	// 预加载Orders、Profile、Role多个关联属性
	// ps: 预加载字段，必须是User的属性
	db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
	//// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
	//// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to
}
