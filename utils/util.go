package utils

import (
	"ginEssential/model"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

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
func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
