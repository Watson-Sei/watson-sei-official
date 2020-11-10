package Controllers

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		db := Config.DBConnect()
		loginUser := user.Username
		loginPassword := user.Password
		db.Find(&Models.User{}, "username =?", loginUser).Scan(&user)
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPassword)); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err, "login": loginUser})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"token": Models.CreateJWTToken(user.Username),
		})
	}
}

// Logout
func LogoutPost(context *gin.Context)  {
	err := Models.BlackListSet(context.MustGet("exp").(int))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{})
}