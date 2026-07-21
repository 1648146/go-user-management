// @title 项目接口文档
// @version 1.0
// @description 服务API说明
// @host 127.0.0.1:8080
// @BasePath /

package main

import (
	"go-user-management/controller"
	"go-user-management/dao"
	_ "go-user-management/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 修改为你本地MySQL密码
	dsn := "root:111222@tcp(127.0.0.1:3306)/go_user_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，请检查密码与数据库是否创建")
	}
	dao.InitDB(db)

	// 初始化Gin路由
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controller.RegisterHandler)
		userGroup.POST("/login", controller.LoginHandler)
	}

	// 启动服务
	err = r.Run(":8080")
	if err != nil {
		panic("服务启动失败")
	}
}
