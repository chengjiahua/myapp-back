package main

import (
	"flag"
	"fmt"
	"myapp-back/config"
	"myapp-back/database"
	"myapp-back/handler"
	"myapp-back/middleware"
	"myapp-back/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 定义命令行参数
	configPath := flag.String("-c", "cfg.json", "配置文件路径")
	flag.Parse()

	config.LoadConfig(*configPath)


	fmt.Println("Config loaded successfully")

	// 初始化数据库
	database.InitDB() // 连接数据库
	defer database.DB.Close()
	model.Migrate() // 自动迁移

	router := gin.Default()
	
	// 注册路由
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	
	// 需要认证的路由示例
	auth := router.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			userID := c.MustGet("userID").(uint)
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	router.Run(":8080")
}