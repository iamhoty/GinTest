package main

import (
	"GinCode/dao"
	"GinCode/models"
	"GinCode/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err) // todo 记录日志
	} else {
		// 日志 连接成功
	}
	// 延迟关闭数据库
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 路由
	r := routers.SetUpRouter()
	r.Run()
}
