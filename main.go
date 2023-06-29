package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {
	db := initDB()
	//defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		phone := ctx.PostForm("phone")
		password := ctx.PostForm("password")

		if len(phone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位"})
			return
		}
		//如果名称没有传则返回一个随机字符串
		if len(name) == 0 {
			name = RandomString(6)
		}
		//判断手机号是否存在
		log.Println("判断手机号是否存在")
		if isPhoneExist(db, phone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已存在，不允许注册"})
			return
		}
		//新建用户
		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&newUser)
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

// 返回随机长度字符串
func RandomString(n int) string {
	letter := []byte("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOKLJHGFDSDSCXVBNNHG")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}

// 查询手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 初始化数据库
func initDB() *gorm.DB {
	//driverName := "mysql"
	//host := "localhost"
	//port := "3306"
	//database := "ginEssential"
	//username := "root"
	//password := "123456"
	//charset := "utf-8"
	//args := fmt.Sprint("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
	//	username,
	//	password,
	//	host,
	//	port,
	//	database,
	//	charset)
	//db, err := gorm.Open(driverName, args)
	args := "root:123456@tcp(127.0.0.1:3306)/ginEssential?charset=utf8mb4&parseTime=True&loc=Local"

	log.Println("链接信息:" + args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Println("fail to connect database:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
