package Routes

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Middleware"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(controller Models.Model) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://localhost", "http://localhost:3000", "https://www.watson-sei.tokyo"},
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
			return CorsChecker(origin, cors.DefaultConfig().AllowOrigins)
		},
	}))
	v1 := router.Group("/v1")
	{
		v1.GET("/article/list", Controllers.GetArticleController{Model: controller}.GetArticle)
		v1.POST("/article/create", Middleware.JWTChecker(), Controllers.Controller{Model: controller}.CreateArticle)
		v1.GET("/article/detail/:id", Controllers.Controller{Model: controller}.GetArticleByID)
		v1.PUT("/article/update/:id", Middleware.JWTChecker(), Controllers.Controller{Model: controller}.UpdateArticle)
		v1.DELETE("/article/delete/:id", Middleware.JWTChecker(), Controllers.Controller{Model: controller}.DeleteArticle)
		v1.GET("/article/tags", Controllers.Controller{Model: controller}.GetAllTag)
		v1.GET("/article/tags/:tag", Controllers.Controller{Model: controller}.GetArticleByTag)

		v1.POST("/upload/image", Middleware.JWTChecker(), Controllers.UploadImage)

	}
	admin := router.Group("/admin")
	{
		// サイトを後悔する際に/signupはコメントアウトして使えないようにする(cli createsuperuserコマンドを作成するのが解決策)
		admin.POST("/signup", Controllers.Controller{Model: controller}.SignupPost)
		admin.POST("/login", Controllers.Controller{Model: controller}.LoginPost)
		admin.GET("/logout", Middleware.JWTChecker(), Controllers.LogoutPost)
		admin.GET("/refresh", Middleware.RefreshChecker(), Controllers.Controller{Model: controller}.RefreshGet)
		admin.GET("/main", Middleware.JWTChecker(), Controllers.Home)
	}
	return router
}

// Cors Checker
func CorsChecker(origin string, corsList []string) bool {
	for _, v := range corsList {
		if origin == v {
			return true
		}
	}
	return false
}