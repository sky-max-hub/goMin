package main

import (
	"flag"
	"fmt"
	"os"
	"zrw/goMin/config"
	"zrw/goMin/db"
	"zrw/goMin/logger"
	"zrw/goMin/redis"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "配置文件路径")
	flag.Parse()

	configPath = "./config/config.json"

	if configPath == "" {
		fmt.Printf("Config Path must be assigned.")
		os.Exit(1)
	}

	var err error
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("Init config failed. Error is %v", err)
		os.Exit(1)
	}

	logConfig := config.GetConfig().LogConfig

	err = logger.InitLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		fmt.Printf("Init logger failed. Error is %v", err)
		os.Exit(1)
	}

	err = db.InitEngine()
	if err != nil {
		fmt.Printf("Init DB failed. Error is %v", err)
		os.Exit(1)
	}

	err = redis.InitRedis()
	if err != nil {
		fmt.Printf("Init Redis failed. Error is %v", err)
		os.Exit(1)
	}

	logger.GetLogger().Info("Init success.")
}
