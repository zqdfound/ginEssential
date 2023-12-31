package common

import (
	"ginEssential/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 初始化数据库
func InitDB() *gorm.DB {

	//todo 从配置文件读取
	database := viper.GetString("datasource.database")
	args := "root:123456@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("链接信息:" + args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Println("fail to connect database:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}
