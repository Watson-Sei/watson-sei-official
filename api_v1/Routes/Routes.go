package Routes

import (
	"github.com/Watson-Sei/jwt-gin/api_v1/middleware"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(cors.New(cors.Config{
	//	// アクセスを許可したいアクセス元
	//	AllowOrigins: []string{
	//		"https://localhost",
	//	},
	//	// 許可したいHTTPメソッドの一覧
	//	AllowMethods: []string{
	//		"POST",
	//		"GET",
	//		"OPTIONS",
	//		"PUT",
	//		"DELETE",
	//	},
	//	// 許可したいHTTPリクエストヘッダーの一覧
	//	AllowHeaders: []string{
	//		"Access-Control-Allow-Credentials",
	//		"Access-Control-Allow-Headers",
	//		"Content-Type",
	//		"Content-Length",
	//		"Accept-Encoding",
	//		"Authorization",
	//	},
	//	// preflightリクエストの結果をキャッシュする時間
	//	MaxAge: 24 * time.Hour,
	//}))
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