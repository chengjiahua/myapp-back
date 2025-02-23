package http

import (
	"log"
	"myapp-back/config"
	"myapp-back/handler"
	"myapp-back/middleware"

	"github.com/gin-gonic/gin"
)

func InitHttp() {

	router := gin.Default()

	loginV1 :=handler.NewLogin()

	// 注册路由
	router.POST("/v1/register", loginV1.Register)
	router.POST("/v1/login", loginV1.Login)

	// 需要认证的路由示例
	auth := router.Group("/v1")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			userID := c.MustGet("userID").(uint)
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	// 启动服务器
	if err := router.Run(config.Cfg.Service.Address); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}