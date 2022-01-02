package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	CreatedAt time.Time // 默认创建时间字段， 在创建时，如果该字段值为零值，则使用当前时间填充
	UpdatedAt int       // 默认更新时间字段， 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	//Updated   int64 `gorm:"autoUpdateTime:nano"` // 自定义字段， 使用时间戳填纳秒数充更新时间
	Updated int64 `gorm:"autoUpdateTime:milli"` //自定义字段， 使用时间戳毫秒数填充更新时间
	Created int64 `gorm:"autoCreateTime"`       //自定义字段， 使用时间戳秒数填充创建时间
}

// TableName 设置表名，可以通过给 struct 类型定义 TableName 函数，返回当前 struct 绑定的 mysql 表名是什么
func (u User) TableName() string {
	//绑定 MYSQL 表名为 users
	return "users"
}

/*
自动生成的表结构如下：
mysql> show create table users \G
*************************** 1. row ***************************
       Table: users
Create Table: CREATE TABLE `users` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `updated` bigint(20) DEFAULT NULL,
  `created` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

*/

func main() {
	//配置MySQL连接参数
	username := "root"       //账号
	password := "123456"     //密码
	host := "192.168.24.147" //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "gorm_test"    //数据库名

	//通过前面的数据库参数，拼接 MYSQL DSN，其实就是数据库连接串（数据源名称）
	//MYSQL dsn 格式：{username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似 {username} 使用花括号包着的名字都是需要替换的参数
	/*
		注意：想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数：https://github.com/go-sql-driver/mysql#parameters)
		要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4 查看 此文章（https://mathiasbynens.be/notes/mysql-utf8mb4） 获取详情
		Sets the location for time.Time values (when using parseTime=true). "Local" sets the system's location.
	*/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	//1.连接 MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	//2.自动建表
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	//3.插入一条用户数据
	//定义一个用户，并初始化数据
	u := User{}
	//下面代码会自动生成 SQL 语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tommy','123456','1540824823')
	if err := db.Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//5.查询并返回第一条数据
	//定义需要保存数据的 struct 变量
	u = User{}
	result := db.First(&u) //自动生成 sql： SELECT * FROM `users`  LIMIT 1
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Printf("%+v", u)
}
