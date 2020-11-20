package Routes

import (
	"github.com/Watson-Sei/jwt-gin/api_v1/middleware"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/user/list", Controllers.GetArticle)
		v1.POST("/user/create", Middleware.JWTChecker(), Controllers.CreateArticle)
		v1.GET("/user/detail/:id", Controllers.GetArticleByID)
		v1.PUT("/user/update/:id", Middleware.JWTChecker(), Controllers.UpdateArticle)
		v1.DELETE("/user/delete/:id", Middleware.JWTChecker(), Controllers.DeleteArticle)
		v1.POST("/upload/image", Controllers.UploadImage)
	}
	admin := router.Group("/admin")
	{
		admin.POST("/signup", Controllers.SignupPost)
		admin.POST("/login", Controllers.LoginPost)
		admin.GET("/logout", Middleware.JWTChecker(), Controllers.LogoutPost)
		admin.GET("/refresh", middleware.RefreshChecker(), Controllers.RefreshGet)
		admin.GET("/main", Middleware.JWTChecker(), Controllers.Main)
	}
	return router
}