package Controllers

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
)

// ユーザー登録
func (c Controller) SignupPost(context *gin.Context)  {
	var user Models.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		err := c.Model.CreateUser(user.Username, user.Password)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"message":"signup success"})
		}
	}
}

// ログイン関数
func (c Controller) LoginPost(context *gin.Context)  {
	var user Models.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		if err := c.Model.PasswordChecker(user.Username, user.Password); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		}
		tokens, err := Models.CreateJWTToken(user.Username, user.ID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		}
		context.JSON(http.StatusOK, tokens)
	}
}

// Logout
func LogoutPost(context *gin.Context)  {
	exp := context.MustGet("exp").(float64)
	token := context.MustGet("accessToken").(string)
	err := Models.BlackListSet(int64(exp), token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		context.Abort()
	} else {
		context.JSON(http.StatusOK, gin.H{"message":"Logout成功しました"})
	}
}

// Token Refresh
func (c Controller) RefreshGet(context *gin.Context)  {
	userId := context.MustGet("userId").(float64)
	refreshToken := context.MustGet("refreshToken").(string)
	exp := context.MustGet("exp").(float64)
	tokens, err := c.Model.RefreshJWTToken(uint(userId), refreshToken, int64(exp))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		context.Abort()
	} else {
		context.JSON(http.StatusOK, tokens)
	}
}