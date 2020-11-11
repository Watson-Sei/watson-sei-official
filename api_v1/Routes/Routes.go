package Routes

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", Controllers.GetArticle)
		v1.POST("/create", Controllers.CreateArticle)
		v1.GET("/:id", Controllers.GetArticleByID)
		v1.PUT("/:id", Controllers.UpdateArticle)
		v1.DELETE("/:id", Controllers.DeleteArticle)
	}
	admin := router.Group("/admin")
	{
		admin.POST("/signup", Controllers.SignupPost)
		admin.POST("/login", Controllers.LoginPost)
		admin.POST("/logout", Middleware.JWTChecker(), Controllers.LogoutPost)
		admin.GET("/main", Middleware.JWTChecker(), Controllers.Main)
	}
	return router
}