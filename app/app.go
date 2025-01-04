package app

import (
	"backend/controllers/common"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/logger"
	"backend/internal/web"
	"backend/models"
)

func Run(configPath string) {
	// 读取配置
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	logger.InitLog("debug", "log/log.log")

	// 初始化数据库
	db.InitDB(config)

	// 初始化casbin
	common.InitCsbinEnforcer()

	// 数据库迁移
	models.Migration()

	// 初始化web服务
	web.InitWeb(config)

	// 启动web服务
	logger.Debug(config.Web.Domain + "站点已启动...")
}
