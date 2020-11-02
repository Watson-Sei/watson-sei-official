package Controllers

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
)

func SignupGet(context *gin.Context) {
	context.HTML(200, "signup.html", gin.H{})
}

func SignupPost(context *gin.Context) {
	var form Models.User
	// バリデーション処理
	if err := context.Bind(&form); err != nil {
		context.HTML(http.StatusBadRequest, "signup.html", gin.H{"err":err})
		context.Abort()
	} else {
		username := context.PostForm("username")
		password := context.PostForm("password")
		// 登録ユーザーが重複していた場合に弾く処理
		err := Models.CreateUser(username, password)
		if err != nil {
			context.HTML(http.StatusBadRequest, "signup.html", gin.H{"err":err})
		}
		context.Redirect(302, "/")
	}
}