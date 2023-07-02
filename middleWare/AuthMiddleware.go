package middleWare

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取header
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		//validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			ctx.Abort()
			return
		}
		//passed
		userId := claims.UserId
		DB := common.InitDB()
		var user model.User
		DB.First(&user, userId)
		//验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			ctx.Abort()
			return
		}
		//将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
