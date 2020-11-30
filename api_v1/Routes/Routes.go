package Routes

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Middleware"

	// "github.com/gin-gonic/contrib/cors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://localhost", "http://localhost:3000","https://www.watson-sei.tokyo"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "https://localhost" || origin == "http://localhost:3000" || origin == "https://www.watson-sei.tokyo" {
				return true
			} else {
				return false
			}
		},
	}))
	v1 := router.Group("/v1")
	{
		v1.GET("/article/list", Controllers.GetArticle)
		v1.POST("/article/create", Middleware.JWTChecker(), Controllers.CreateArticle)
		v1.GET("/article/detail/:id", Controllers.GetArticleByID)
		v1.PUT("/article/update/:id", Middleware.JWTChecker(), Controllers.UpdateArticle)
		v1.DELETE("/article/delete/:id", Middleware.JWTChecker(), Controllers.DeleteArticle)
		v1.GET("/article/tags", Controllers.GetAllTag)
		v1.GET("/article/tags/:tag", Controllers.GetArticleByTag)

		v1.POST("/upload/image", Controllers.UploadImage)

	}
	admin := router.Group("/admin")
	{
		admin.POST("/signup", Controllers.SignupPost)
		admin.POST("/login", Controllers.LoginPost)
		admin.GET("/logout", Middleware.JWTChecker(), Controllers.LogoutPost)
		admin.GET("/refresh", Middleware.RefreshChecker(), Controllers.RefreshGet)
		admin.GET("/main", Middleware.JWTChecker(), Controllers.Main)
	}
	return router
}
