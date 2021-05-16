package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_project?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 测试连通性 ping通则返回nil
	err = DB.DB().Ping()
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return nil
}

func Close() {
	// 延迟关闭数据库
	DB.Close()
}
