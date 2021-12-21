package main

import "gorm.io/gorm"

/*
CREATE TABLE `food` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID，商品Id',
  `name` varchar(30) NOT NULL COMMENT '商品名',
  `price` decimal(10,2) unsigned  NOT NULL COMMENT '商品价格',
  `type_id` int(10) unsigned NOT NULL COMMENT '商品类型Id',
  `createtime` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',
   PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

// 默认 gorm 对s truct 字段名使用 Snake Case 命名风格转换成 mysql 表字段名(需要转换成小写字母)。
// 提示：Snake Case命名风格，就是各个单词之间用下划线（_）分隔，例如： CreateTime的Snake Case风格命名为create_time

// Food 字段注释说明了gorm 库把 struct 字段转换为表字段名长什么样子。
type Food struct {
	Id     int     //表字段名为：id
	Name   string  //表字段名为：name
	Price  float64 //表字段名为：price
	TypeId int     //表字段名为：type_id
	//字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是 golang 的基础语法，不同的库会定义不同的标签，有不同的含义
	//gorm标签语法：gorm:"标签定义"
	//标签定义部分，多个标签定义可以使用分号（;）分隔。如 gorm:"column:id; PRIMARY_KEY"
	CreateTime int64 `gorm:"column:createtime"` //表字段名为：createtime
}

type User struct {
	gorm.Model // 嵌入gorm.Model的字段
	Name       string
}

func main() {

}
