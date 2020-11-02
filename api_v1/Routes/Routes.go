package Routes

import (
	"time"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://www.watson-sei.tokyo",
			"https://watson-sei.tokyo",
			"https://localhost",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge: 24 * time.Hour,
	}))
	{
		v1.GET("/", Controllers.GetArticle)
		v1.POST("/create", Controllers.CreateArticle)
		v1.GET("/:id", Controllers.GetArticleByID)
		v1.PUT("/:id", Controllers.UpdateArticle)
		v1.DELETE("/:id", Controllers.DeleteArticle)
	}

	return router
}