package routers

import (
	"GinCode/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	// 结构体与数据库映射
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static") // 项目的相对路径
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	// v1
	v1Group := r.Group("/v1") // todo
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.GET("/todo/:id", func(c *gin.Context) { // 查看某一条

		})
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
