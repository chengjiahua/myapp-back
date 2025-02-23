package main

import (
	"flag"
	"fmt"
	"myapp-back/config"
	"myapp-back/database"
	"myapp-back/http"
	"myapp-back/model"
)

func main() {
	// 定义命令行参数
	configPath := flag.String("c", "./cfg.json", "配置文件路径")
	flag.Parse()

	config.LoadConfig(*configPath) // 加载配置文件


	fmt.Println("Config loaded successfully")

	// 初始化数据库
	database.InitDB() // 连接数据库
	defer database.DB.Close()
	model.Migrate() // 自动迁移

	http.InitHttp() // 初始化 HTTP 服务器
	
}