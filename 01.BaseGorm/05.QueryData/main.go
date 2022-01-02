package main

import (
	"fmt"
	"gorm_test/tools"
	"time"
)

//商品
type Food struct {
	Id    int
	Title string
	Price float32
	Stock int
	Type  int
	//mysql datetime, date 类型字段，可以和 golang time.Time 类型绑定，详细说明请参考：gorm 连接数据库章节。
	CreateTime time.Time
}

//为 Food 绑定表名
func (v Food) TableName() string {
	return "foods"
}

func main() {
	//获取DB
	db := tools.GetDB()
	defer tools.Close()

	/*------------------------------------------------------query-----------------------------------------------------*/
	// 1.Take：查询一条记录
	//等价于：SELECT * FROM `foods`   LIMIT 1
	food := Food{}
	db.Take(&food)
	fmt.Printf("1. %+v \n", food)

	//2.First：查询一条记录，根据主键 ID 排序（正序），返回第一条记录
	//等价于：SELECT * FROM `foods` ORDER BY `foods`.`id` ASC LIMIT 1
	food = Food{}
	db.First(&food) // 注意：food 变量中的字段要都为零值，否则就会被用作 Where 条件
	fmt.Printf("2. %+v \n", food)

	//3.Last：查询一条记录，根据主键 ID 排序（倒序），返回第一条记录
	//等价于：SELECT * FROM `foods` ORDER BY `foods`.`id` DESC LIMIT 1
	food = Food{}
	db.Last(&food) // 注意：food 变量中的字段要都为零值，否则就会被用作 Where 条件
	fmt.Printf("3. %+v \n", food)

	//4.Find：查询多条记录，Find 函数返回的是一个数组
	//等价于：SELECT * FROM `foods`
	foods := []Food{}
	db.Find(&foods)
	fmt.Printf("4. %+v \n", foods)

	//5.Pluck：查询一列值
	//等价于：SELECT title FROM `foods`
	titles := []string{}
	//db.Table("foods").Pluck("title", &titles)
	db.Model(&Food{}).Pluck("title", &titles)
	fmt.Printf("5. %+v \n", titles)

	/*------------------------------------------------------where-----------------------------------------------------*/
	//等价于: SELECT * FROM `foods`  WHERE (id = '2') LIMIT 1
	food = Food{}
	db.Where("id = ?", 2).Take(&food)
	fmt.Printf("6. %+v \n", food)

	//in 语句
	//等价于: SELECT * FROM `foods` WHERE (id in ('1','2','5','6')) LIMIT 1
	food = Food{}
	db.Where("id in (?)", []int{2, 5, 6}).Take(&food)
	fmt.Printf("7. %+v \n", food)

	//等价于: SELECT * FROM `foods`  WHERE (create_time >= '2018-11-06 00:00:00' and create_time <= '2018-11-06 23:59:59')
	//这里使用了两个问号(?)占位符，后面传递了两个参数替换两个问号。
	foods = []Food{}
	db.Where("create_time >= ? and create_time <= ?", "2022-01-02 09:12:18", "2022-01-02 09:13:55").Find(&foods)
	fmt.Printf("8. %+v \n", foods)

	//like 语句
	//等价于: SELECT * FROM `foods` WHERE (title like '%可乐%')
	foods = []Food{}
	db.Where("title like ?", "%A%").Find(&foods)
	fmt.Printf("9. %+v \n", foods)

	/*------------------------------------------------------select----------------------------------------------------*/

	//等价于: SELECT id,title FROM `foods` WHERE (id = '1') LIMIT 1
	food = Food{}
	db.Select("id,title").Where("id = ?", 1).Take(&food)
	fmt.Printf("10. %+v \n", food)

	//这种写法是直接往 Select 函数传递数组，数组元素代表需要选择的字段名
	food = Food{}
	db.Select([]string{"id", "title"}).Where("id = ?", 1).Take(&food)
	fmt.Printf("11. %+v \n", food)

	//可以直接书写聚合语句
	//等价于: SELECT count(*) as total FROM `foods`
	total := []int{}
	db.Model(&Food{}).Select("count(*) as total").Pluck("total", &total)
	fmt.Printf("12. %+v \n", total)

	/*------------------------------------------------------order----------------------------------------------------*/
	//等价于: SELECT * FROM `foods`  WHERE (create_time >= '2018-11-06 00:00:00') ORDER BY create_time desc
	foods = []Food{}
	db.Where("create_time >= ?", "2022-01-02 09:12:18").Order("create_time desc").Find(&foods)
	fmt.Printf("13. %+v \n", foods)

	/*--------------------------------------------------limit&offset--------------------------------------------------*/
	//等价于: SELECT * FROM `foods` ORDER BY create_time desc LIMIT 10 OFFSET 0
	foods = []Food{}
	db.Order("create_time desc").Limit(10).Offset(0).Find(&foods)
	fmt.Printf("14. %+v \n", foods)

	/*-----------------------------------------------------count------------------------------------------------------*/
	var totalT int64 = 0
	//等价于: SELECT count(*) FROM `foods`
	db.Model(Food{}).Count(&totalT)
	fmt.Printf("15. %+v \n", totalT)

	/*----------------------------------------------------group by----------------------------------------------------*/
	//统计每个商品分类下面有多少个商品
	//定一个 Result 结构体类型，用来保存查询结果
	type Result struct {
		Type  int
		Total int
	}

	results := []Result{}
	//等价于: SELECT type, count(*) as  total FROM `foods` GROUP BY type HAVING (total > 0)
	db.Model(Food{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)
	fmt.Printf("16. %+v \n", results)

	/*----------------------------------------------直接执行 sql 语句-------------------------------------------------*/
	//例子:
	sql := "SELECT type, count(*) as total FROM `foods` where create_time > ? GROUP BY type HAVING (total > 0)"

	//因为sql语句使用了一个问号(?)作为绑定参数,
	//所以需要传递一个绑定参数(Raw 第二个参数).
	//Raw 函数支持绑定多个参数
	results = []Result{}
	db.Raw(sql, "2022-01-02 09:12:18").Scan(&results)
	fmt.Printf("17. %+v \n", results)
}
