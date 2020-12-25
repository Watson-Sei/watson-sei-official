
package Routes

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(Controller Controllers.Controller) *gin.Engine {
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
		v1.GET("/article/list", Controller.GetArticle)
		v1.POST("/article/create", Middleware.JWTChecker(), Controller.CreateArticle)
		v1.GET("/article/detail/:id", Controller.GetArticleByID)
		v1.PUT("/article/update/:id", Middleware.JWTChecker(), Controller.UpdateArticle)
		v1.DELETE("/article/delete/:id", Middleware.JWTChecker(), Controller.DeleteArticle)
		v1.GET("/article/tags", Controller.GetAllTag)
		v1.GET("/article/tags/:tag", Controller.GetArticleByTag)

		v1.POST("/upload/image", Middleware.JWTChecker(), Controllers.UploadImage)

	}
	admin := router.Group("/admin")
	{
		// サイトを後悔する際に/signupはコメントアウトして使えないようにする(cli createsuperuserコマンドを作成するのが解決策)
		admin.POST("/signup", Controller.SignupPost)
		admin.POST("/login", Controller.LoginPost)
		admin.GET("/logout", Middleware.JWTChecker(), Controllers.LogoutPost)
		admin.GET("/refresh", Middleware.RefreshChecker(), Controller.RefreshGet)
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
