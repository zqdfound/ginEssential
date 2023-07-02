package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	InitConfig() //初始化配置
	r := gin.Default()
	r = CollectRoute(r)
	r.Run()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("加载配置失败")
		panic(any(err))
	}
}
