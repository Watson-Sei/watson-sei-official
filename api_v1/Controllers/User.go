package Controllers

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
)

// ユーザー登録
func SignupPost(context *gin.Context)  {
	var user Models.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		err := Models.CreateUser(user.Username, user.Password)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"message":"signup success"})
		}
	}
}

// ログイン関数
func LoginPost(context *gin.Context)  {
	var user Models.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		if err := Models.PasswordChecker(user.Username, user.Password); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"token":Models.CreateJWTToken(user.Username),
		})
	}
}

// Logout
func LogoutPost(context *gin.Context)  {
	exp := context.MustGet("exp").(float64)
	token := context.MustGet("token").(string)
	err := Models.BlackListSet(int64(exp), token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		context.Abort()
	} else {
		context.JSON(http.StatusOK, gin.H{"message":"Logout成功しました"})
	}
}